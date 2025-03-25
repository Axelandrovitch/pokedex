package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationsApiData struct {
	Count       int        `json:"count"`
	Results     []Location `json:"results"`
	NextURL     string     `json:"next"`
	PreviousURL string     `json:"previous"`
	BaseURL     string
	FirstFectch bool
}

func FetchLocations(url string) (LocationsApiData, error) {
	var apiResponse LocationsApiData
	res, err := http.Get(url)
	if err != nil {
		return LocationsApiData{}, fmt.Errorf("failed to fetch data %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return LocationsApiData{}, fmt.Errorf("response failed with status code %d ", res.StatusCode)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationsApiData{}, fmt.Errorf("failed to read response body %v", err)
	}
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return LocationsApiData{}, fmt.Errorf("could not unmarshal JSON %w", err)
	}
	return apiResponse, nil
}

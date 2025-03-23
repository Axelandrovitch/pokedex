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
	res, err := http.Get(url)
	if err != nil {
		return LocationsApiData{}, fmt.Errorf("Failed to fetch data %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return LocationsApiData{}, fmt.Errorf("Response failed with status code %d ", res.StatusCode)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationsApiData{}, fmt.Errorf("Failed to read response body %v", err)
	}
	var apiResponse LocationsApiData
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return LocationsApiData{}, fmt.Errorf("Could not unmarshal JSON %w", err)
	}
	return apiResponse, nil
}

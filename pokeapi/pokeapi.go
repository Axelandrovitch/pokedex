package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationsApiData struct {
	Count       int        `json:"count"`
	Results     []Location `json:"results"`
	NextUrl     string     `json:"next"`
	PreviousUrl string     `json:"previous"`
	BaseUrl     string
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
		return LocationsApiData{}, fmt.Errorf("Could not unmarshal JSON %d", err)
	}
	return apiResponse, nil
}

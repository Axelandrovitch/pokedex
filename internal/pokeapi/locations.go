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
	BaseLocationAreaURL     string
	FirstFectch bool
}

func (client *Client)FetchLocations(url string) (LocationsApiData, error) {
	var apiResponse LocationsApiData
	if val, ok := client.cache.Get(url); ok {
		err := json.Unmarshal(val, &apiResponse)
		if err != nil {
			return LocationsApiData{}, fmt.Errorf("could not unmarshal JSON from cache %w", err)
		}
		return apiResponse, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationsApiData{}, err
	}
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return LocationsApiData{}, fmt.Errorf("failed to fetch data from %s: %v", url, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationsApiData{}, fmt.Errorf("failed to read response body %v", err)
	}
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return LocationsApiData{}, fmt.Errorf("could not unmarshal JSON %w", err)
	}
	client.cache.Add(url, body)
	return apiResponse, nil
}

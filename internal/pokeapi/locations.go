package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Axelandrovitch/pokedex/internal/pokecache"
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

type Client struct {
	cache		pokecache.Cache
	httpClient	http.Client
	baseURL		string
}

func NewClient(baseURL string, timeout, cacheDuration time.Duration) error 

func FetchLocations(cache *pokecache.Cache, url string) (LocationsApiData, error) {
	var apiResponse LocationsApiData
	if val, ok := cache.Get(url); ok {
		err := json.Unmarshal(val, &apiResponse)
		if err != nil {
			return LocationsApiData{}, fmt.Errorf("could not unmarshal JSON %w", err)
		}
		return apiResponse, nil
	}
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
	cache.Add(url, body)
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return LocationsApiData{}, fmt.Errorf("could not unmarshal JSON %w", err)
	}
	return apiResponse, nil
}

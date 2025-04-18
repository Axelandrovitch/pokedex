package main

import (
	"fmt"
)

func commandMap(config *Config) error {
	LocationsApiData := config.Locations
	var url string
	if LocationsApiData.FirstFectch {
		LocationsApiData.FirstFectch = false
		url = LocationsApiData.BaseLocationAreaURL
	} else {
		if LocationsApiData.NextURL == "" {
			fmt.Println("No more locations to explore in this direction!")
			return nil
		}
		url = LocationsApiData.NextURL
	}
	UpdatedApiData, err := config.Client.FetchLocations(url)
	if err != nil {
		return fmt.Errorf("failed to update API data %w", err)
	}
	LocationsApiData.Results = UpdatedApiData.Results
	LocationsApiData.NextURL = UpdatedApiData.NextURL
	LocationsApiData.PreviousURL = UpdatedApiData.PreviousURL
	locations := LocationsApiData.Results
	for _, location := range locations {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapBack(config *Config) error {
	LocationsApiData := config.Locations
	if LocationsApiData.PreviousURL == "" {
		fmt.Println("No more locations to explore in this direction!")
		return nil
	}
	updatedApiData, err := config.Client.FetchLocations(LocationsApiData.PreviousURL)
	if err != nil {
		return fmt.Errorf("failed to update API data %w", err)
	}
	*LocationsApiData = updatedApiData
	locations := LocationsApiData.Results
	for _, location := range locations {
		fmt.Println(location.Name)
	}
	return nil
}

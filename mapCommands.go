package main

import (
	"fmt"

	"github.com/Axelandrovitch/pokedex/internal/pokeapi"
)

func commandMap(config *Config) error {
	LocationsApiData := config.Locations
	if LocationsApiData.FirstFectch {
		UpdatedApiData, err := pokeapi.FetchLocations(config.Cache, LocationsApiData.BaseURL)
		if err != nil {
			return fmt.Errorf("failed to update API data %w", err)
		}
		*LocationsApiData = UpdatedApiData
		LocationsApiData.FirstFectch = false
		locations := LocationsApiData.Results
		for _, location := range locations {
			fmt.Println(location.Name)
		}
		return nil
	}
	if LocationsApiData.NextURL == "" {
		fmt.Println("No more locations to explore in this direction!")
		return nil
	}
	UpdatedApiData, err := pokeapi.FetchLocations(config.Cache, LocationsApiData.NextURL)
	if err != nil {
		return fmt.Errorf("failed to update API data %w", err)
	}
	*LocationsApiData = UpdatedApiData
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
	updatedApiData, err := pokeapi.FetchLocations(config.Cache, LocationsApiData.PreviousURL)
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

package main

import (
	"fmt"

	"github.com/Axelandrovitch/pokedex/internal/pokeapi"
)

func validLocation(userLocation string, locations []pokeapi.Location) (string, bool) {

	for _, location := range locations {
		if location.Name == userLocation {
			return location.URL, true
		}
	}
	return "", false
}

func commandExplore(config *Config) error {
	fmt.Println("Explore !")
	for _, cmd := range config.cmdArgs {
		fmt.Println(cmd)
	}
	if config.Locations.FirstFectch {
		fmt.Println("You need to map location areas first!")
		return nil
	}
	userLocation := config.cmdArgs[0]
	location, exists := validLocation(userLocation, config.Locations.Results)
	if !exists {
		return fmt.Errorf("given location: %s is invalid", userLocation)
	}
	locationURL := config.Locations.CurrentURL + "/" + location
	pokemons, err := config.Client.FetchPokemons(locationURL)
	if err != nil {
		return fmt.Errorf("failed to fetch pokemons from given location %w", err)
	}
	for _, pokemon := range pokemons {
		fmt.Println(pokemon)
	}
	return nil
}

package main

import (
	"fmt"

	"github.com/Axelandrovitch/pokedex/internal/pokeapi"
)

// func validLocation(userLocation string, locations []pokeapi.Location) (string, bool) {

// 	for _, location := range locations {
// 		if location.Name == userLocation {
// 			return location.URL, true
// 		}
// 	}
// 	return "", false
// }

func printPokemons(pokemons []pokeapi.PokemonEncounter) {
	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemons {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
}

func commandExplore(config *Config) error {
	// if config.Locations.FirstFectch {
	// 	fmt.Println("You need to map location areas first!")
	// 	return nil
	// }
	userLocation := config.cmdArgs[0]
	// locationURL, exists := validLocation(userLocation, config.Locations.Results)
	// if !exists {
	// 	return fmt.Errorf("given location: %s is invalid", userLocation)
	// }
	locationURL := config.Locations.BaseLocationAreaURL + userLocation
	fmt.Printf("Exploring %s...\n", userLocation)
	pokemonEncounters, err := config.Client.FetchPokemons(locationURL)
	if err != nil {
		return fmt.Errorf("failed to fetch pokemons from given location %w", err)
	}
	printPokemons(pokemonEncounters)
	return nil
}

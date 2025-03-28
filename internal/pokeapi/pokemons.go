package pokeapi

import (
	"encoding/json"
	"fmt"
)

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonEncounters struct {
	Pokemon        Pokemon          `json:"pokemon"`
}

func (client *Client) FetchPokemons(url string) (PokemonEncounters, error) {
	var pokemonEncounters PokemonEncounters
	if val, ok := client.cache.Get(url); ok {
		err := json.Unmarshal(val, &pokemonEncounters)
		if err != nil {
			return PokemonEncounters{}, fmt.Errorf("could not unmarshal JSON from cache %w", err)
		}
		return pokemonEncounters, nil
	}
	
	return pokemonEncounters, nil
}

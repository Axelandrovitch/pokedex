package pokeapi

import (
	"encoding/json"
	"fmt"
)

type Pokemons []string

func (client *Client) FetchPokemons(url string) (Pokemons, error) {
	var pokemons Pokemons
	if val, ok := client.cache.Get(url); ok {
		err := json.Unmarshal(val, &pokemons)
		if err != nil {
			return Pokemons{}, fmt.Errorf("could not unmarshal JSON from cache %w", err)
		}
		return pokemons, nil
	}
	return pokemons, nil
}

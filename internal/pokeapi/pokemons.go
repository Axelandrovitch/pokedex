package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

func (client *Client) FetchPokemons(url string) ([]PokemonEncounter, error) {
	var APIResponse struct {
		PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
	}
	if val, ok := client.cache.Get(url); ok {
		err := json.Unmarshal(val, &APIResponse)
		if err != nil {
			return nil, fmt.Errorf("could not unmarshal JSON from cache %w", err)
		}
		return APIResponse.PokemonEncounters, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from %s: %v", url, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response from body %v", err)
	}
	err = json.Unmarshal(body, &APIResponse)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON %w", err)
	}
	client.cache.Add(url, body)
	return APIResponse.PokemonEncounters, nil
}

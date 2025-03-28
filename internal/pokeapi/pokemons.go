package pokeapi

import (
	// "fmt"
)

type Pokemons []string

func validLocation(userLocation string, locations []Location) (string, bool) {

	for _, location := range locations {
		if location.Name == userLocation {
			return location.URL, true
		}
	}
	return "", false
}

func fetchPokemons(arguments []string, locations []Location) (Pokemons, error) {
	// chosenLocation := arguments[0]
	// // url, ok := validLocation(chosenLocation, locations)
	// if !ok {
	// 	return nil, fmt.Errorf("given location could not be found in current map")
	// }
	// var pokemons Pokemons
	// // req, err := http.NewRequest("GET", url, nil)
	// // if err != nil {
	// // 	return nil, err
	// // }
	return nil, nil
}

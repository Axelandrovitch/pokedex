package pokeapi

type Pokemons []string

func fetchPokemons(arguments []string, url string) (Pokemons, error) {
	var pokemons Pokemons
	chosenLocation := arguments[0]
	newURL := url + chosenLocation

	return nil, nil
}

package main

import (
	"fmt"

	"github.com/Axelandrovitch/pokedex/internal/pokeapi"
	"github.com/Axelandrovitch/pokedex/internal/pokecache"
)

type Config struct {
	Cache     	*pokecache.Cache
	Locations 	*pokeapi.LocationsApiData
	currentCmd	string
	cmdArgv		[]string
}

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

func main() {
	fmt.Println("Welcome to the Pokedex!")
	readFromStdin()
}



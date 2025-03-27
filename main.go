package main

import (
	"fmt"
	"time"

	"github.com/Axelandrovitch/pokedex/internal/pokeapi"
	"github.com/Axelandrovitch/pokedex/internal/pokecache"
)

type Config struct {
	Cache      *pokecache.Cache
	Locations  *pokeapi.LocationsApiData
	currentCmd string
	cmdArgs    []string
	currentURL string
}

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

func main() {
	client := pokeapi.NewClient(5 * time.Second, 5 * time.Minute)
	fmt.Println("Welcome to the Pokedex!")
	readFromStdin()
}

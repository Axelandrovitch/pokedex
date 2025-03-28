package main

import (
	"fmt"
	"time"

	"github.com/Axelandrovitch/pokedex/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient("https://pokeapi.co/api/v2/", 5*time.Second, 5*time.Minute)
	fmt.Println("Welcome to the Pokedex!")
	readFromStdin(client)
}

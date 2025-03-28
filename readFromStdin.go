package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Axelandrovitch/pokedex/internal/pokeapi"
)

type Config struct {
	Client     *pokeapi.Client
	Locations  *pokeapi.LocationsApiData
	currentCmd string
	// currentURL string
	cmdArgs    []string
}

func cleanInput(text string) []string {
	if len(text) == 0 {
		return []string{}
	}
	trimmed := strings.TrimSpace(strings.ToLower(text))
	split := strings.Fields(trimmed)
	return split
}

func readFromStdin(client pokeapi.Client) {
	scanner := bufio.NewScanner(os.Stdin)

	config := &Config{
		Client: &client,
		Locations: &pokeapi.LocationsApiData{
			BaseLocationAreaURL:     "https://pokeapi.co/api/v2/location-area/",
			FirstFectch: true,
		},
	}
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		executeUserCommand(input, config)
	}
}

func executeUserCommand(input []string, config *Config) {
	if cmd, validCommand := getCommands()[input[0]]; validCommand {
		config.currentCmd = input[0]
		config.cmdArgs = nil
		if len(input) > 1 {
			config.cmdArgs = input[1:]
		}
		cmd.Callback(config)
	} else {
		fmt.Println("Unknown command")
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/Axelandrovitch/pokedex/internal/pokeapi"
	"github.com/Axelandrovitch/pokedex/internal/pokecache"
)

func cleanInput(text string) []string {
	if len(text) == 0 {
		return []string{}
	}
	trimmed := strings.TrimSpace(strings.ToLower(text))
	split := strings.Fields(trimmed)
	return split
}

func readFromStdin() {
	scanner := bufio.NewScanner(os.Stdin)
	config := &Config{
		Cache: pokecache.NewCache(time.Minute * 5),
		Locations: &pokeapi.LocationsApiData{
			FirstFectch: true,
			BaseURL:     "https://pokeapi.co/api/v2/location-area/",
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
		command := input[0]
		executeUserCommand(command, config)
	}
}

func executeUserCommand(command string, config *Config) {
	if cmd, validCommand := getCommands()[command]; validCommand {
		cmd.Callback(config)
	} else {
		fmt.Println("Unknown command")
	}
}
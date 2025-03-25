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

type Config struct {
	Cache     *pokecache.Cache
	Locations *pokeapi.LocationsApiData
}

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

var supportedCommands map[string]cliCommand

func initCommands() {
	supportedCommands = map[string]cliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"help": {
			Name:        "help",
			Description: "Lists available commands",
			Callback:    commandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Displays location areas in the Pokemon world",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays location areas in the Pokemon world",
			Callback:    commandMapBack,
		},
	}
}

func main() {

	initCommands()
	fmt.Println("Welcome to the Pokedex!")
	readFromStdin()
}

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
	if cmd, validCommand := supportedCommands[command]; validCommand {
		cmd.Callback(config)
	} else {
		fmt.Println("Unknown command")
	}
}

func commandExit(config *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *Config) error {
	fmt.Println("Usage: ")
	for _, cmd := range supportedCommands {
		fmt.Printf(" %s: %s\n", cmd.Name, cmd.Description)
	}
	return nil
}

func commandMap(config *Config) error {
	LocationsApiData := config.Locations
	if LocationsApiData.FirstFectch {
		UpdatedApiData, err := pokeapi.FetchLocations(config.Cache, LocationsApiData.BaseURL)
		if err != nil {
			return fmt.Errorf("failed to update API data %w", err)
		}
		*LocationsApiData = UpdatedApiData
		LocationsApiData.FirstFectch = false
		locations := LocationsApiData.Results
		for _, location := range locations {
			fmt.Println(location.Name)
		}
		return nil
	}
	if LocationsApiData.NextURL == "" {
		fmt.Println("No more locations to explore in this direction!")
		return nil
	}
	UpdatedApiData, err := pokeapi.FetchLocations(config.Cache, LocationsApiData.NextURL)
	if err != nil {
		return fmt.Errorf("failed to update API data %w", err)
	}
	*LocationsApiData = UpdatedApiData
	locations := LocationsApiData.Results
	for _, location := range locations {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapBack(config *Config) error {
	LocationsApiData := config.Locations
	if LocationsApiData.PreviousURL == "" {
		fmt.Println("No more locations to explore in this direction!")
		return nil
	}
	updatedApiData, err := pokeapi.FetchLocations(config.Cache, LocationsApiData.PreviousURL)
	if err != nil {
		return fmt.Errorf("failed to update API data %w", err)
	}
	*LocationsApiData = updatedApiData
	locations := LocationsApiData.Results
	for _, location := range locations {
		fmt.Println(location.Name)
	}
	return nil
}

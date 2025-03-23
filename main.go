package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Axelandrovitch/pokedex/pokeapi"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*pokeapi.LocationsApiData) error
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
	var apiData pokeapi.LocationsApiData
	apiData.FirstFectch = true
	apiData.BaseUrl = "https://pokeapi.co/api/v2/location-area/"
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
		executeUserCommand(command, &apiData)
	}
}

func executeUserCommand(command string, apiData *pokeapi.LocationsApiData) {
	if cmd, validCommand := supportedCommands[command]; validCommand {
		cmd.Callback(apiData)
	} else {
		fmt.Println("Unknown command")
	}
}

func commandExit(apiData *pokeapi.LocationsApiData) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("Error exiting the program")
}

func commandHelp(apiData *pokeapi.LocationsApiData) error {
	fmt.Println("Usage: ")
	for _, cmd := range supportedCommands {
		fmt.Printf(" %s: %s\n", cmd.Name, cmd.Description)
	}
	return fmt.Errorf("Error executing help command")
}

func commandMap(apiData *pokeapi.LocationsApiData) error {
	if *&apiData.FirstFectch == true {
		UpdatedApiData, err := pokeapi.FetchLocations(apiData.BaseUrl)
		if err != nil {
			return fmt.Errorf("Failed to update API data %d", err)
		}
		*apiData = UpdatedApiData
		*&apiData.FirstFectch = false
		locations := apiData.Results
		for _, location := range locations {
			fmt.Println(location.Name)
		}
		return nil
	}
	if apiData.NextUrl == "" {
		return fmt.Errorf("No more locations to explore in this direction!")
	}
	UpdatedApiData, err := pokeapi.FetchLocations(apiData.NextUrl)
	if err != nil {
		return fmt.Errorf("Failed to update API data %d", err)
	}
	*apiData = UpdatedApiData
	locations := apiData.Results
	for _, location := range locations {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapBack(apiData *pokeapi.LocationsApiData) error {

	if apiData.PreviousUrl != "" {
		UpdatedApiData, err := pokeapi.FetchLocations(apiData.PreviousUrl)
		if err != nil {
			return fmt.Errorf("Failed to update API data %d", err)
		}
		*apiData = UpdatedApiData
	}
	locations := apiData.Results
	for _, location := range locations {
		fmt.Println(location.Name)
	}
	return nil
}

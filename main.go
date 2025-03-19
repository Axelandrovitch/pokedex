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
	Callback    func() error
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
	}
}

func main() {
	initCommands()
	fmt.Println("Welcome to the Pokedex!")
	//	readFromStdin()
	commandMap()
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
		executeUserCommand(command)
	}
}

func executeUserCommand(command string) {
	if cmd, validCommand := supportedCommands[command]; validCommand {
		cmd.Callback()
	} else {
		fmt.Println("Unknown command")
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("Error exiting the program")
}

func commandHelp() error {
	fmt.Println("Usage: ")
	for _, cmd := range supportedCommands {
		fmt.Printf(" %s: %s\n", cmd.Name, cmd.Description)
	}
	return fmt.Errorf("Error executing help command")
}

func commandMap() error {
	apiData, err := pokeapi.FetchLocations()
	if err != nil {
		return err
	}
	locations := apiData.Results
	for _, location := range locations {
		fmt.Println(location.Name)
	}
	return nil
}

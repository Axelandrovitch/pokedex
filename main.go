package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	}
}

func main() {
	initCommands()
	fmt.Println(readFromStdin())
}

func cleanInput(text string) []string {
	if len(text) == 0 {
		return []string{}
	}
	trimmed := strings.TrimSpace(strings.ToLower(text))
	split := strings.Fields(trimmed)

	return split
}

func readFromStdin() []string {
	for {
		fmt.Print("Pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input[0]) > 0 {
			command := input[0]
			executeUserCommand(command)
		}
	}
}

func executeUserCommand(command string) error {
	if cmd, validCommand := supportedCommands[command]; validCommand {
		return cmd.Callback()
	}
	fmt.Println("Unknown command")
	return fmt.Errorf("Unknown command")
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return fmt.Errorf("Error exiting the program")
}

func commandHelp() error {
	fmt.Println("Usage <.pokedex> <command>")
	fmt.Println("Available commands:")
	for _, cmd := range supportedCommands {
		fmt.Println(cmd.Name)
		fmt.Printf("\t %s\n", cmd.Description)
	}
	return fmt.Errorf("Error executing help command")
}

package main

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
		"explore": {
			Name:		"explore",
			Description: "Lists all the Pokémon listed in the a given location",
			Callback: commandExplore,
		},
	}
}
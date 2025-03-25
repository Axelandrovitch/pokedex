package main

import "fmt"

func commandHelp(config *Config) error {
	fmt.Println("Usage: ")
	for _, cmd := range getCommands(){
		fmt.Printf(" %s: %s\n", cmd.Name, cmd.Description)
	}
	return nil
}
package main

import "fmt"

func commandExplore(config *Config) error {
	fmt.Println("Explore !")
	for _, cmd := range config.cmdArgs {
		fmt.Println(cmd)
	}
	if !config.Locations.FirstFectch {
		//look for pokemons

	}
	return nil
}

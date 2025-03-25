package main

import "fmt"

func commandExplore(c *Config) error {
	fmt.Println("Explore !")
	for _, cmd := range c.cmdArgv {
		fmt.Println(cmd)
	}
	if !c.Locations.FirstFectch {
		//look for pokemons
	}
	return nil
}

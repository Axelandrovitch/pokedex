package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

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
		fmt.Print("Pokedex >")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := cleanInput(scanner.Text())
		fmt.Printf("Your command was: %s\n", input[0])
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		input.Scan()

		command := cleanInput(input.Text())
		if len(command) > 0 {
			fmt.Printf("Your command was: %s\n", command[0])
		}
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	trimmed := strings.TrimSpace(lower)
	inputs := strings.Fields(trimmed)
	return inputs
}

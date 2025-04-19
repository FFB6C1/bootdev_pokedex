package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func repl(cfg config) {
	input := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	for {
		fmt.Print("Pokedex > ")
		input.Scan()

		command := cleanInput(input.Text())
		if len(command) > 0 {
			if comm, ok := commands[command[0]]; ok {
				comm.callback(&cfg)
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	trimmed := strings.TrimSpace(lower)
	inputs := strings.Fields(trimmed)
	return inputs
}

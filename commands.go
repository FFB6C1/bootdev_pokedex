package main

import (
	"fmt"
	"os"
)

type command struct {
	name     string
	desc     string
	callback func() error
}

func GetCommands() map[string]command {
	commands := map[string]command{
		"help": {
			name:     "help",
			desc:     "Displays a list of available commands.",
			callback: commandHelp,
		},
		"exit": {
			name:     "exit",
			desc:     "Exits the program.",
			callback: commandExit,
		},
	}
	return commands
}

func commandHelp() error {
	commands := GetCommands()
	fmt.Print("List of Commands:\n\n")
	for _, command := range commands {
		fmt.Printf("  %s: %s \n", command.name, command.desc)
	}
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

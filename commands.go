package main

import (
	"fmt"
	"os"
)

type command struct {
	name     string
	desc     string
	callback func(config) error
}

func GetCommandsList() []command {
	commandsList := []command{
		{
			name:     "help",
			desc:     "Displays a list of available commands.",
			callback: commandHelp,
		},
		{
			name:     "exit",
			desc:     "Exits the program.",
			callback: commandExit,
		},
		{
			name:     "map",
			desc:     "Displays the next page of locations.",
			callback: commandMap,
		},
		{
			name:     "mapb",
			desc:     "Displays the previous page of locations.",
			callback: commandMapB,
		},
	}

	return commandsList
}

func GetCommands() map[string]command {
	commandsList := GetCommandsList()

	commands := map[string]command{}
	for _, item := range commandsList {
		commands[item.name] = item
	}

	return commands
}

func commandHelp(_config config) error {
	commands := GetCommandsList()
	fmt.Print("List of Commands:\n\n")
	for _, command := range commands {
		fmt.Printf("  %s: %s \n", command.name, command.desc)
	}
	return nil
}

func commandExit(_config config) error {
	os.Exit(0)
	return nil
}

func commandMap(_config config) error {
	//url := config.mapNext
	//data, err := apiInteraction.handleRequestAndResponse(url)
	return nil
}

func commandMapB(_config config) error {
	return nil
}

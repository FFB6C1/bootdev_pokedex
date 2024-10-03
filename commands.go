package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/FFB6C1/bootdev_pokedex/internal/apiInteraction"
)

type command struct {
	name     string
	desc     string
	callback func(*config) error
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

func commandHelp(_config *config) error {
	commands := GetCommandsList()
	fmt.Print("List of Commands:\n\n")
	for _, command := range commands {
		fmt.Printf("  %s: %s \n", command.name, command.desc)
	}
	return nil
}

func commandExit(_config *config) error {
	os.Exit(0)
	return nil
}

func commandMap(config *config) error {
	err := mapMove(config)
	if err != nil {
		return err
	}
	config.mapOffset += 20
	config.mapStep = true
	return nil
}

func commandMapB(config *config) error {
	if config.mapStep {
		config.mapOffset -= 40
	} else {
		config.mapOffset -= 20
	}

	if config.mapOffset < 0 {
		config.mapOffset = 0
		return fmt.Errorf("Cannot go back from here. Returned to start.")
	}

	err := mapMove(config)
	if err != nil {
		return err
	}
	return nil
}

func mapMove(config *config) error {
	url := config.mapAPI + "offset=" + strconv.Itoa(config.mapOffset) + "limit=" + strconv.Itoa(config.mapLimit)
	data, err := apiInteraction.LocationRequest(url, &config.cache)
	if err != nil {
		return err
	}
	for _, location := range data.Results {
		fmt.Println(location.Name)
	}
	return nil
}

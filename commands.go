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
	callback func(*config, ...string) error
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
		{
			name:     "explore",
			desc:     "Explores a location. Usage: explore [area name].",
			callback: commandExplore,
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

// Command Functions

func commandHelp(_ *config, _ ...string) error {
	commands := GetCommandsList()
	fmt.Print("List of Commands:\n\n")
	for _, command := range commands {
		fmt.Printf("  %s: %s \n", command.name, command.desc)
	}
	return nil
}

func commandExit(_ *config, _ ...string) error {
	os.Exit(0)
	return nil
}

func commandMap(config *config, _ ...string) error {
	err := mapMove(config)
	if err != nil {
		return err
	}
	config.mapOffset += 20
	config.mapStep = true
	return nil
}

func commandMapB(config *config, _ ...string) error {
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

func commandExplore(config *config, params ...string) error {
	if len(params) == 0 {
		fmt.Println("Where would you like to explore?")
		fmt.Println("Use 'explore' followed by an area name, or use the help command.")
		return nil
	} else {
		fmt.Println("Exploring: " + params[0] + "...")
	}
	url := config.mapAPI + params[0]
	data, err := apiInteraction.AreaRequest(url, &config.cache)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println("Found Pokemon:")

	for _, pokemon := range data.PokemonEncounters {
		fmt.Println("- " + pokemon.Pokemon.Name)
	}

	return nil
}

// Helper Functions

func mapMove(config *config) error {
	url := config.mapAPI + "?offset=" + strconv.Itoa(config.mapOffset) + "limit=" + strconv.Itoa(config.mapLimit)
	data, err := apiInteraction.LocationRequest(url, &config.cache)
	if err != nil {
		return err
	}
	for _, location := range data.Results {
		fmt.Println(location.Name)
	}
	return nil
}

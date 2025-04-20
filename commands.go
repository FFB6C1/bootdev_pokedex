package main

import (
	pokeapi "bootdev_pokedex/internal/apiInteraction"
	"fmt"
	"net/url"
	"os"
)

type config struct {
	next     string
	previous string
	client   pokeapi.Client
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	var commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 locations in the Pokemon World",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations in the Pokemon World",
			callback:    commandMapB,
		},
	}
	return commands
}

func commandExit(_ *config, _ ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(_ *config, _ ...string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	commands := getCommands()
	for _, c := range commands {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	return nil
}

func commandMap(cfg *config, _ ...string) error {
	data, err := pokeapi.LocationGet(cfg.next, cfg.client)
	if err != nil {
		handleError("commandMap", err, false)
	}

	helperUpdateNextPrevious(cfg, data)

	for _, l := range data.Results {
		fmt.Println(l.Name)
	}

	return nil
}

func commandMapB(cfg *config, _ ...string) error {
	if cfg.previous == "" {
		fmt.Println("You're on the first page!")
		return nil
	}

	data, err := pokeapi.LocationGet(cfg.previous, cfg.client)
	if err != nil {
		handleError("commandMapB", err, false)
	}

	helperUpdateNextPrevious(cfg, data)

	for _, l := range data.Results {
		fmt.Println(l.Name)
	}
	return nil
}

func helperUpdateNextPrevious(cfg *config, data pokeapi.Location) {
	cfg.next = helperGetQuery(data.Next)
	if data.Previous != nil {
		cfg.previous = helperGetQuery(*data.Previous)
	} else {
		cfg.previous = ""
	}
}

func helperGetQuery(fullURL string) string {
	parsed, err := url.Parse(fullURL)
	if err != nil {
		handleError("helperURLSlicer", err, false)
	}
	return string("?" + parsed.RawQuery)
}

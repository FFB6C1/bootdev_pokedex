package main

import (
	pokeapi "bootdev_pokedex/internal/apiInteraction"
	"fmt"
	"math/rand/v2"
	"net/url"
	"os"
)

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
		"explore": {
			name:        "explore",
			description: "Displays a list of the pokemon available in an area. Usage: 'explore [area]",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a wild pokemon! Usage: 'catch [pokemon-name]'",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a previously caught pokemon. Usage: 'inspect [pokemon-name]'",
			callback:    commandInspect,
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

func commandExplore(cfg *config, commands ...string) error {
	if len(commands) == 0 {
		fmt.Println("Please choose a place to explore. Use 'map' to see some possible locations.\nUsage: 'explore [areaname]'")
		return nil
	}
	locationPokemon, err := pokeapi.LocationExploreGet(commands[0], cfg.client)
	if err != nil {
		fmt.Println("Could not find that area! Did you spell its name correctly?")
		return nil
	}
	fmt.Printf("Exploring %s...\n", commands[0])
	fmt.Println("Found Pokemon:")
	for _, pokemon := range locationPokemon.PokemonEncounters {
		fmt.Println(" - " + pokemon.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *config, commands ...string) error {
	if len(commands) == 0 {
		fmt.Println("Please choose a pokemon to try to catch! Try exploring to see what's around.")
		fmt.Println("Usage: catch [pokemon-name]")
	}
	pokemon, err := pokeapi.PokemonGet(commands[0], cfg.client)
	if err != nil {
		fmt.Println("Could not find that pokemon! Did you spell its name correctly?")
		return nil
	}
	fmt.Println("Throwing a Pokeball at " + commands[0] + "...")
	catch := helperCatchTry(pokemon.BaseExperience)
	if catch {
		cfg.pokedex[commands[0]] = pokemon
		fmt.Println("You caught " + commands[0] + "!")
	} else {
		fmt.Println("Oh no! It got away...")
	}
	return nil
}

func commandInspect(cfg *config, commands ...string) error {
	if len(commands) == 0 {
		fmt.Println("Please choose a pokemon to inspect! Remember that this command only works with caught pokemon.")
		fmt.Println("Usage: 'inspect <pokemon-name>")
		return nil
	}
	pokemon, ok := cfg.pokedex[commands[0]]
	if !ok {
		fmt.Println("you have not caught that pokemon!")
		return nil
	}
	helperPrintInformation(pokemon)
	return nil
}

// Helper Functions below here.

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

func helperCatchTry(difficulty int) bool {
	if difficulty > 400 {
		difficulty = 399
	}
	number := rand.IntN(400)
	return number > difficulty
}

func helperPrintInformation(pokemon pokeapi.Pokemon) {
	fmt.Println("Name: " + pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, elem := range pokemon.Types {
		fmt.Printf("  -%s\n", elem.Type.Name)
	}
}

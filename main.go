package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var cliCommands = map[string]cliCommand{
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
}

var helpText string = "\nWelcome to the Pokedex!\nCommands:\n\n"

func main() {
	for _, item := range cliCommands {
		helpText = helpText + item.name + ": " + item.description + "\n"
	}
	replLoop()
}

func replLoop() {
	running := true
	commandScanner := bufio.NewScanner(os.Stdin)
	for running {
		fmt.Print("Pokedex > ")

		commandScanner.Scan()
		err := commandScanner.Err()
		if err != nil {
			log.Fatal(err)
		}

		command := strings.ToLower(commandScanner.Text())
		if _, ok := cliCommands[command]; ok {
			cliCommands[command].callback()
		} else {
			fmt.Println("\nUse 'help' to see a list of valid commands!")
		}
	}
}

func commandHelp() error {
	fmt.Println(helpText)
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//setup
	commandMap := GetCommands()
	scanner := bufio.NewScanner(os.Stdin)
	config := initiateConfig()
	fmt.Println("Welcome to the Pokedex! Please type a command to continue, or type 'help' for a list of commands.")

	//REPL
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		if command, ok := commandMap[input]; ok {
			command.callback(config)
		} else {
			fmt.Println("Invalid command. Type 'help' for a list of commands.")
		}
	}
}

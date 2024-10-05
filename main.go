package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
		input := strings.Split(strings.ToLower(scanner.Text()), " ")
		if command, ok := commandMap[input[0]]; ok {
			err := command.callback(&config, input[1:]...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Invalid command. Type 'help' for a list of commands.")
		}
	}
}

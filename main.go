package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/FFB6C1/bootdev_pokedex/internal/pokecache"
)

func main() {
	//setup
	commandMap := GetCommands()
	scanner := bufio.NewScanner(os.Stdin)
	cache := pokecache.NewCache(7 * time.Second)
	config := initiateConfig(cache)
	fmt.Println("Welcome to the Pokedex! Please type a command to continue, or type 'help' for a list of commands.")

	//REPL
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		if command, ok := commandMap[input]; ok {
			err := command.callback(&config)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Invalid command. Type 'help' for a list of commands.")
		}
	}
}

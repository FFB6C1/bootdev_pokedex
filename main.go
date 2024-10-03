package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
		input := strings.Split(scanner.Text(), " ")
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

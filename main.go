package main

import (
	pokeapi "bootdev_pokedex/internal/apiInteraction"
	"time"
)

func main() {
	client := pokeapi.GetClient(2 * time.Second)
	config := config{
		next:     "?offset=0&limit=20",
		previous: "",
		client:   client,
	}
	repl(config)
}

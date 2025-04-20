package main

import (
	pokeapi "bootdev_pokedex/internal/apiInteraction"
	"time"
)

type config struct {
	next     string
	previous string
	client   pokeapi.Client
	pokedex  map[string]pokeapi.Pokemon
}

func main() {
	client := pokeapi.GetClient(2*time.Second, 5*time.Second)
	pokedex := map[string]pokeapi.Pokemon{}
	config := config{
		next:     "?offset=0&limit=20",
		previous: "",
		client:   client,
		pokedex:  pokedex,
	}
	repl(config)
}

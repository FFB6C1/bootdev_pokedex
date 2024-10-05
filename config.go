package main

import (
	"time"

	"github.com/FFB6C1/bootdev_pokedex/internal/pokecache"

	"github.com/FFB6C1/bootdev_pokedex/internal/pokedex"
)

type config struct {
	cache      pokecache.Cache
	pokedex    pokedex.Pokedex
	mapAPI     string
	pokemonAPI string
	mapOffset  int
	mapLimit   int
	mapStep    bool // set to 'true' by map and 'false' by mapb, used for backtracking
}

func initiateConfig() config {
	return config{
		cache:      pokecache.NewCache(7 * time.Second),
		pokedex:    pokedex.NewPokedex(),
		mapAPI:     "/location-area/",
		pokemonAPI: "/pokemon/",
		mapOffset:  0,
		mapLimit:   10,
		mapStep:    false,
	}
}

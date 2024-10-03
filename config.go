package main

import (
	"github.com/FFB6C1/bootdev_pokedex/internal/pokecache"
)

type config struct {
	cache     pokecache.Cache
	mapAPI    string
	mapOffset int
	mapLimit  int
	mapStep   bool // set to 'true' by map and 'false' by mapb, used for backtracking
}

func initiateConfig(cache pokecache.Cache) config {
	return config{
		cache:     cache,
		mapAPI:    "/location-area/",
		mapOffset: 0,
		mapLimit:  20,
		mapStep:   false,
	}
}

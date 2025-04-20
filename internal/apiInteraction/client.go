package pokeapi

import (
	"bootdev_pokedex/internal/pokecache"
	"net/http"
	"time"
)

type Client struct {
	client http.Client
	cache  *pokecache.Cache
}

func GetClient(timeout, cacheDuration time.Duration) Client {
	client := http.Client{
		Timeout: timeout,
	}
	cache := pokecache.NewCache(cacheDuration)
	return Client{
		client: client,
		cache:  cache,
	}
}

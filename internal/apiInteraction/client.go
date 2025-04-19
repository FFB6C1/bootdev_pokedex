package pokeapi

import (
	"net/http"
	"time"
)

type Client struct {
	client http.Client
}

func GetClient(timeout time.Duration) Client {
	c := http.Client{
		Timeout: timeout,
	}
	return Client{
		client: c,
	}
}

package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type Location struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func apiGet(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	return body, nil
}

func LocationGet(url string, client Client) (Location, error) {

	fullURL := baseURL + locationURL + url
	locationJson, err := apiGet(fullURL)
	if err != nil {
		return Location{}, err
	}

	locations := Location{}

	if err = json.Unmarshal(locationJson, &locations); err != nil {
		return Location{}, err
	}

	return locations, nil
}

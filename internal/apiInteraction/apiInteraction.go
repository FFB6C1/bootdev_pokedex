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

type LocationPokemon struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func apiGet(url string, client Client) ([]byte, error) {
	if bytes, ok := client.cache.Get(url); ok {
		return bytes, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	client.cache.Add(url, body)
	return body, nil
}

func LocationGet(position string, client Client) (Location, error) {

	fullURL := baseURL + locationURL + position
	locationJson, err := apiGet(fullURL, client)
	if err != nil {
		return Location{}, err
	}

	locations := Location{}

	if err = json.Unmarshal(locationJson, &locations); err != nil {
		return Location{}, err
	}

	return locations, nil
}

func LocationExploreGet(area string, client Client) (LocationPokemon, error) {
	fullURL := baseURL + locationURL + area
	locationPokemonJson, err := apiGet(fullURL, client)
	if err != nil {
		return LocationPokemon{}, err
	}

	locationPokemon := LocationPokemon{}

	if err = json.Unmarshal(locationPokemonJson, &locationPokemon); err != nil {
		return LocationPokemon{}, err
	}

	return locationPokemon, nil
}

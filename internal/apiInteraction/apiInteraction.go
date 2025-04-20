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

type Pokemon struct {
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`
	BaseExperience int `json:"base_experience"`
	Forms          []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	GameIndices []struct {
		GameIndex int `json:"game_index"`
		Version   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version"`
	} `json:"game_indices"`
	Height    int `json:"height"`
	HeldItems []struct {
		Item struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"item"`
		VersionDetails []struct {
			Rarity  int `json:"rarity"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"held_items"`
	ID                     int    `json:"id"`
	IsDefault              bool   `json:"is_default"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt  int `json:"level_learned_at"`
			MoveLearnMethod struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move_learn_method"`
			Order        any `json:"order"`
			VersionGroup struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version_group"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Name    string `json:"name"`
	Order   int    `json:"order"`
	Species struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
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

func PokemonGet(name string, client Client) (Pokemon, error) {
	fullURL := baseURL + pokemonURL + name
	pokemonJson, err := apiGet(fullURL, client)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}

	if err = json.Unmarshal(pokemonJson, &pokemon); err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}

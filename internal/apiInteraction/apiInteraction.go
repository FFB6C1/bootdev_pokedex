package apiInteraction

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/FFB6C1/bootdev_pokedex/internal/pokecache"
	"github.com/FFB6C1/bootdev_pokedex/internal/pokedex"
)

func LocationRequest(url string, cache *pokecache.Cache) (Location, error) {
	if data, ok := checkCache(url, cache); ok {
		unmarshalledData := Location{}
		err := json.Unmarshal(data, &unmarshalledData)
		if err != nil {
			return Location{}, err
		}
		return unmarshalledData, nil
	}

	req, err := makeRequest(baseURL + url)
	if err != nil {
		return Location{}, err
	}

	res, err := getResponse(req)
	if err != nil {
		return Location{}, err
	}

	defer res.Body.Close()

	data, err := readResponse(res)
	if err != nil {
		return Location{}, err
	}

	cache.Add(url, data)

	unmarshalledData := Location{}
	if err := json.Unmarshal(data, &unmarshalledData); err != nil {
		return Location{}, err
	}

	return unmarshalledData, nil
}

func AreaRequest(url string, cache *pokecache.Cache) (Area, error) {
	if data, ok := checkCache(url, cache); ok {
		unmarshalledData := Area{}
		err := json.Unmarshal(data, &unmarshalledData)
		if err != nil {
			return Area{}, err
		}
		return unmarshalledData, nil
	}

	req, err := makeRequest(baseURL + url)
	if err != nil {
		fmt.Println("request error")
		return Area{}, err
	}

	res, err := getResponse(req)
	if err != nil {
		fmt.Println("response error")
		return Area{}, err
	}
	defer res.Body.Close()

	data, err := readResponse(res)
	if err != nil {
		return Area{}, err
	}

	cache.Add(url, data)

	unmarshalledData := Area{}
	if err := json.Unmarshal(data, &unmarshalledData); err != nil {
		newErr := fmt.Errorf("Can't find pokemon! Make sure you spelled the area correctly. To find areas, use the map command!")
		return Area{}, newErr
	}

	return unmarshalledData, nil
}

func PokemonRequest(url string, cache *pokecache.Cache) (pokedex.Pokemon, error) {
	if data, ok := checkCache(url, cache); ok {
		unmarshalledData := pokedex.Pokemon{}
		err := json.Unmarshal(data, &unmarshalledData)
		if err != nil {
			return pokedex.Pokemon{}, err
		}
		return unmarshalledData, nil
	}

	req, err := makeRequest(baseURL + url)
	if err != nil {
		fmt.Println("Request Error")
		return pokedex.Pokemon{}, err
	}

	res, err := getResponse(req)
	if err != nil {
		fmt.Println("Response Error")
		return pokedex.Pokemon{}, err
	}
	defer res.Body.Close()

	data, err := readResponse(res)
	cache.Add(url, data)

	unmarshalledData := pokedex.Pokemon{}
	if err := json.Unmarshal(data, &unmarshalledData); err != nil {
		newError := fmt.Errorf("Could not read pokemon data - did you use a valid pokemon name?")
		return pokedex.Pokemon{}, newError
	}

	return unmarshalledData, nil
}

func checkCache(key string, cache *pokecache.Cache) ([]byte, bool) {
	data, ok := cache.Get(key)
	return data, ok
}

func makeRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	return req, err
}

func getResponse(request *http.Request) (*http.Response, error) {
	client := http.Client{}
	res, err := client.Do(request)
	return res, err
}

func readResponse(res *http.Response) ([]byte, error) {
	data, err := io.ReadAll(res.Body)
	return data, err
}

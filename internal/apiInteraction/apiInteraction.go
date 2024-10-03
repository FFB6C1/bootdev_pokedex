package apiInteraction

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/FFB6C1/bootdev_pokedex/internal/pokecache"
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

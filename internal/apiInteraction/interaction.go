package apiInteraction

import (
	"encoding/json"
	"net/http"
)

func handleRequestAndResponse(url string) (Location, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return Location{}, err
	}

	defer res.Body.Close()

	decodeStruct := Location{}

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&decodeStruct); err != nil {
		return Location{}, err
	}

	return decodeStruct, nil

}

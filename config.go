package main

type config struct {
	mapNext string
	mapPrev string
}

func initiateConfig() config {
	return config{
		mapNext: "https://pokeapi.co/api/v2/location-area/?offset=0limit=20",
		mapPrev: "",
	}
}

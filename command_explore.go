package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandExploreTea(conf *config, inputs ...string) string {
	if len(inputs) < 1 {
		return "Please provide a location area to explore"
	}
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", inputs[0])
	var areaInfo LocationArea
	val, found := conf.Cache.Get(url)
	if found {
		err := json.Unmarshal(val, &areaInfo)
		if err != nil {
			return err.Error()
		}
	} else {
		res, err := http.Get(url)
		if err != nil {
			return err.Error()
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err.Error()
		}
		conf.Cache.Add(url, body)
		err = json.Unmarshal(body, &areaInfo)
		if err != nil {
			return err.Error()
		}
	}
	output := fmt.Sprintf("Exploring %s...\nFound Pokemon:\n", inputs[0])
	for _, encounter := range areaInfo.PokemonEncounters {
		output = fmt.Sprintf("%s- %s\n", output, encounter.Pokemon.Name)
	}
	return output
}

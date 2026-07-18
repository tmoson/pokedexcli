package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func commandExplore(conf *config, inputs ...string) error {
	if len(inputs) < 1 {
		return errors.New("Please provide a location area to explore")
	}
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", inputs[0])
	var areaInfo LocationArea
	val, found := conf.cache.Get(url)
	if found {
		err := json.Unmarshal(val, &areaInfo)
		if err != nil {
			return err
		}
	} else {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		conf.cache.Add(url, body)
		err = json.Unmarshal(body, &areaInfo)
		if err != nil {
			return err
		}
	}
	fmt.Printf("Exploring %s...\nFound Pokemon:\n", inputs[0])
	for _, encounter := range areaInfo.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}
	return nil
}

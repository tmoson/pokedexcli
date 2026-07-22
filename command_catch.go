package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

func commandCatch(conf *config, inputs ...string) error {
	if len(inputs) == 0 {
		return errors.New("Can't catch nothing!")
	}
	pokemon := inputs[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemon)
	var pokemonInfo Pokemon
	val, cached := conf.cache.Get(url)
	info, caught := conf.pokedex[pokemon]
	if cached {
		err := json.Unmarshal(val, &pokemonInfo)
		if err != nil {
			return err
		}
	} else if caught {
		pokemonInfo = info
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
		err = json.Unmarshal(body, &pokemonInfo)
		if err != nil {
			return err
		}
	}
	// chose 400 because experience looks to max out at 306,
	// and this will give difficult (mostly legendary) pokemon an ~23% catch rate
	catchRate := 400 - pokemonInfo.BaseExperience
	if rand.Intn(400) < catchRate {
		fmt.Printf("%s was caught!\n", pokemon)
		conf.pokedex[pokemon] = pokemonInfo
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
	}
	return nil
}
func commandCatchTea(conf *config, inputs ...string) string {
	if len(inputs) == 0 {
		return "Can't catch nothing!"
	}
	pokemon := inputs[0]
	output := fmt.Sprintf("Throwing a Pokeball at %s...\n", pokemon)
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemon)
	var pokemonInfo Pokemon
	val, cached := conf.cache.Get(url)
	info, caught := conf.pokedex[pokemon]
	if cached {
		err := json.Unmarshal(val, &pokemonInfo)
		if err != nil {
			return err.Error()
		}
	} else if caught {
		pokemonInfo = info
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
		conf.cache.Add(url, body)
		err = json.Unmarshal(body, &pokemonInfo)
		if err != nil {
			return err.Error()
		}
	}
	// chose 400 because experience looks to max out at 306,
	// and this will give difficult (mostly legendary) pokemon an ~23% catch rate
	catchRate := 400 - pokemonInfo.BaseExperience
	if rand.Intn(400) < catchRate {
		conf.pokedex[pokemon] = pokemonInfo
		output += fmt.Sprintf("%s was caught!\n", pokemon)
	} else {
		output += fmt.Sprintf("%s escaped!\n", pokemon)
	}
	return output
}

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
)

func commandCatch(conf *config, inputs ...string) error {
	if len(inputs) == 0 {
		return errors.New("Can't catch nothing!")
	}
	pokemon := inputs[0]
	pokeball := "Pokeball"
	if len(inputs) > 1 {
		switch strings.ToLower(inputs[1]) {
		case "great":
			pokeball = "Great Ball"
		case "ultra":
			pokeball = "Ultra Ball"
		default:
			return errors.New("Only have Pokeballs, Great Balls, and Ultra Balls. Specify 'Great' or 'Ultra' when catching to use great or ultra balls")
		}
	}
	fmt.Printf("Throwing %s at %s...\n", pokeball, pokemon)
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemon)
	var pokemonInfo Pokemon
	val, cached := conf.cache.Get(url)
	if cached {
		err := json.Unmarshal(val, &pokemonInfo)
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
		err = json.Unmarshal(body, &pokemonInfo)
		if err != nil {
			return err
		}
	}
	// chose 400 because experience looks to max out at 306,
	// and this will give difficult (mostly legendary) pokemon an ~23% base catch rate
	catchRate := 400 - pokemonInfo.BaseExperience
	switch pokeball {
	case "Great Ball":
		catchRate += int(catchRate / 2)
	case "Ultra Ball":
		catchRate *= 2
	}
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
	pokeball := "Pokeball"
	if len(inputs) > 1 {
		switch strings.ToLower(inputs[1]) {
		case "great":
			pokeball = "Great Ball"
		case "ultra":
			pokeball = "Ultra Ball"
		default:
			return "Only have Pokeballs, Great Balls, and Ultra Balls. Specify 'Great' or 'Ultra' when catching to use great or ultra balls"
		}
	}
	output := fmt.Sprintf("Throwing %s at %s...\n", pokeball, pokemon)
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemon)
	var pokemonInfo Pokemon
	val, cached := conf.cache.Get(url)
	if cached {
		err := json.Unmarshal(val, &pokemonInfo)
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
		conf.cache.Add(url, body)
		err = json.Unmarshal(body, &pokemonInfo)
		if err != nil {
			return err.Error()
		}
	}
	// chose 400 because experience looks to max out at 306,
	// and this will give difficult (mostly legendary) pokemon an ~23% basecatch rate
	catchRate := 400 - pokemonInfo.BaseExperience
	switch pokeball {
	case "Great Ball":
		catchRate += int(catchRate / 2)
	case "Ultra Ball":
		catchRate *= 2
	}
	if rand.Intn(400) < catchRate {
		conf.pokedex[pokemon] = pokemonInfo
		output += fmt.Sprintf("%s was caught!\n", pokemon)
	} else {
		output += fmt.Sprintf("%s escaped!\n", pokemon)
	}
	return output
}

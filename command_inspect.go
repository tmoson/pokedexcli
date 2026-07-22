package main

import (
	"errors"
	"fmt"
)

func commandInspect(conf *config, inputs ...string) error {
	if len(inputs) == 0 {
		return errors.New("Need a pokemon to inspect!")
	}
	pokemon := inputs[0]
	info, caught := conf.pokedex[pokemon]
	if !caught {
		fmt.Printf("you have not caught %s!\n", pokemon)
	} else {
		formattedInfo := fmt.Sprintf("Name: %s\nHeight: %d\nWeight: %d\nStats:\n", info.Name, info.Height, info.Weight)
		for _, stat := range info.Stats {
			formattedInfo = fmt.Sprintf("%s -%s: %d\n", formattedInfo, stat.Stat.Name, stat.BaseStat)
		}
		formattedInfo = fmt.Sprintf("%sTypes:\n", formattedInfo)
		for _, pokeType := range info.Types {
			formattedInfo = fmt.Sprintf("%s  - %s\n", formattedInfo, pokeType.Type.Name)
		}
		fmt.Print(formattedInfo)
	}
	return nil
}

func commandInspectTea(conf *config, inputs ...string) string {
	if len(inputs) == 0 {
		return "Need a pokemon to inspect!"
	}
	pokemon := inputs[0]
	info, caught := conf.pokedex[pokemon]
	if !caught {
		return fmt.Sprintf("you have not caught %s!\n", pokemon)
	} else {
		formattedInfo := fmt.Sprintf("Name: %s\nHeight: %d\nWeight: %d\nStats:\n", info.Name, info.Height, info.Weight)
		for _, stat := range info.Stats {
			formattedInfo = fmt.Sprintf("%s -%s: %d\n", formattedInfo, stat.Stat.Name, stat.BaseStat)
		}
		formattedInfo = fmt.Sprintf("%sTypes:\n", formattedInfo)
		for _, pokeType := range info.Types {
			formattedInfo = fmt.Sprintf("%s  - %s\n", formattedInfo, pokeType.Type.Name)
		}
		return formattedInfo
	}
}

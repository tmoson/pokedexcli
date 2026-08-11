package main

import (
	"github.com/tmoson/pokedexcli/internal/pokecache"
	"regexp"
	"strings"
)

type teaCommand struct {
	name        string
	description string
	callback    func(*config, ...string) string
}

func getCommandsTea() map[string]teaCommand {
	return map[string]teaCommand{
		"map": {
			name:        "map",
			description: "Displays the next 20 map locations",
			callback:    commandMapTea,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 map locations",
			callback:    commandMapbTea,
		},
		"explore": {
			name:        "explore",
			description: "Displays all of the encounterable pokemon in a location area",
			callback:    commandExploreTea,
		},
		"catch": {
			name:        "catch",
			description: "Throws a pokeball at the provided pokemon",
			callback:    commandCatchTea,
		},
		"inspect": {
			name:        "catch",
			description: "Inspect a pokemon you've caught",
			callback:    commandInspectTea,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Display all the pokemon in your pokedex",
			callback:    nil, // need to figure out how to properly null out unneeded callbacks
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelpTea,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    nil, // need to figure out how to properly null out unneeded callbacks
		},
	}
}

type config struct {
	nextLocation     string
	previousLocation string
	cache            pokecache.Cache
	pokedex          map[string]Pokemon
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(strings.TrimSpace(text))
	if lowerText == "" {
		return nil
	}
	re := regexp.MustCompile(`\s+`)
	return re.Split(lowerText, -1)
}

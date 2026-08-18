package main

import (
	"github.com/tmoson/pokedexcli/internal/pokecache"
	"regexp"
	"strings"
	"sync"
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
			name:        "inspect",
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
		"save": {
			name:        "save",
			description: "Save progress to provided file name or save0, located in ~/.pokecliSaves",
			callback:    commandSaveTea,
		},
		"load": {
			name:        "load",
			description: "Load file provided or save0 by default, located in ~/.pokecliSaves",
			callback:    commandLoadTea,
		},
	}
}

type config struct {
	NextLocation     string
	PreviousLocation string
	Cache            pokecache.Cache
	Pokedex          map[string]Pokemon
	lock             sync.Mutex
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(strings.TrimSpace(text))
	if lowerText == "" {
		return nil
	}
	re := regexp.MustCompile(`\s+`)
	return re.Split(lowerText, -1)
}

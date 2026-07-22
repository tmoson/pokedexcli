package main

import (
	"bufio"
	"fmt"
	"github.com/tmoson/pokedexcli/internal/pokecache"
	"os"
	"regexp"
	"strings"
	"time"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

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
			callback:    commandPokedexTea,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelpTea,
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

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name:        "map",
			description: "Displays the next 20 map locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 map locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Displays all of the encounterable pokemon in a location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Throws a pokeball at the provided pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "catch",
			description: "Inspect a pokemon you've caught",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Display all the pokemon in your pokedex",
			callback:    commandPokedex,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func repl() {
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	configuration := config{
		nextLocation: "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		cache:        pokecache.NewCache(5 * time.Second),
		pokedex:      make(map[string]Pokemon),
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		instruction := cleanInput(scanner.Text())
		action, ok := commands[instruction[0]]
		if !ok {
			fmt.Printf("Unrecognized command: %v\n", instruction[0])
			continue
		}
		err := action.callback(&configuration, instruction[1:]...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	locationOffset int
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
		locationOffset: 0,
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		command := cleanInput(scanner.Text())[0]
		action, ok := commands[command]
		if !ok {
			fmt.Printf("Unrecognized command: %v\n", command)
			continue
		}
		err := action.callback(&configuration)
		if err != nil {
			fmt.Println(err)
		}
	}
}

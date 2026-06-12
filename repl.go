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
	callback    func() error
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
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		command := cleanInput(scanner.Text())[0]
		action, ok := commands[command]
		if !ok {
			fmt.Printf("Unrecognized command: %v\n", command)
			continue
		}
		action.callback()
	}
}

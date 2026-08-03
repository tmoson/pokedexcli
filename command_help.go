package main

import (
	"fmt"
	"strings"
)

func commandHelp(conf *config, inputs ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandHelpTea(conf *config, inputs ...string) string {
	helpWidth := 70 // no point in calculating this when the help text is static at compile time
	output := "\nWelcome to the Pokedex!\nUsage:\n\n"
	for _, cmd := range getCommandsTea() {
		nameLength := len(cmd.name)
		nameDescriptionSpacing := 9 - nameLength
		remainingSpacing := helpWidth - (nameDescriptionSpacing + nameLength + len(cmd.description))
		output = fmt.Sprintf(
			"%s%s:%s%s%s\n",
			output,
			cmd.name,
			strings.Repeat(" ", nameDescriptionSpacing),
			cmd.description,
			strings.Repeat(" ", remainingSpacing),
		)
	}
	output += "\n"
	return output
}

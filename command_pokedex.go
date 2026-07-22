package main

import (
	"errors"
	"fmt"
)

func commandPokedex(conf *config, inputs ...string) error {
	if len(inputs) > 0 {
		return errors.New("Pokedex doesn't take any input")
	}
	if len(conf.pokedex) == 0 {
		fmt.Println("You haven't caught any pokemon yet!")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for name, _ := range conf.pokedex {
		fmt.Printf(" - %s\n", name)
	}
	fmt.Println("All of these can be inspected with the 'inspect' command")
	return nil
}

func commandPokedexTea(conf *config, inputs ...string) string {
	if len(inputs) > 0 {
		return "Pokedex doesn't take any input"
	}
	if len(conf.pokedex) == 0 {
		return "You haven't caught any pokemon yet!"
	}
	output := "Your Pokedex:\n"
	for name, _ := range conf.pokedex {
		output = fmt.Sprintf("%s - %s\n", output, name)
	}
	output += "All of these can be inspected with the 'inspect' command"
	return output
}

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

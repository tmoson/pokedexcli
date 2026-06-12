package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		command := scanner.Text()
		fmt.Printf("Your command was: %v\n", cleanInput(command)[0])
	}
}

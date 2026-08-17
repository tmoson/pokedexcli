package main

import (
	"log"

	tea "charm.land/bubbletea/v2"
)

func main() {
	m := InitialModel()
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	p := tea.NewProgram(&m)
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

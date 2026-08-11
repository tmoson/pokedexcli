package main

import (
	"charm.land/lipgloss/v2"
	"fmt"
)

func commandInspectTea(conf *config, inputs ...string) string {
	if len(inputs) == 0 {
		return "Need a pokemon to inspect!"
	}
	pokemon := inputs[0]
	info, caught := conf.pokedex[pokemon]
	if !caught {
		return fmt.Sprintf("you have not caught %s!\n", pokemon)
	}

	headerStyle := lipgloss.NewStyle().
		Bold(true).
		BorderForeground(lipgloss.Color("36")).
		BorderStyle(lipgloss.DoubleBorder()).
		Padding(0, 1).
		MarginBottom(1)

	header := headerStyle.Render(fmt.Sprintf("Name: %s", info.Name))

	list := []string{
		fmt.Sprintf("Height: %d M", info.Height/10),
		fmt.Sprintf("Weight: %d kg", info.Weight/100),
		"Stats:",
	}
	for _, stat := range info.Stats {
		list = append(list, fmt.Sprintf("  - %s: %d", stat.Stat.Name, stat.BaseStat))
	}

	list = append(list, "Types:")
	for _, pokeType := range info.Types {
		list = append(list, fmt.Sprintf("  - %s", pokeType.Type.Name))
	}

	return lipgloss.JoinVertical(lipgloss.Left, header, lipgloss.JoinVertical(lipgloss.Left, list...))
}

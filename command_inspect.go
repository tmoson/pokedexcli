package main

import (
	"errors"
	"fmt"
	"charm.land/lipgloss/v2"
)

func commandInspect(conf *config, inputs ...string) error {
	if len(inputs) == 0 {
		return errors.New("Need a pokemon to inspect!")
	}
	pokemon := inputs[0]
	info, caught := conf.pokedex[pokemon]
	if !caught {
		fmt.Printf("you have not caught %s!\n", pokemon)
	} else {
		formattedInfo := fmt.Sprintf("Name: %s\nHeight: %d\nWeight: %d\nStats:\n", info.Name, info.Height, info.Weight)
		for _, stat := range info.Stats {
			formattedInfo = fmt.Sprintf("%s -%s: %d\n", formattedInfo, stat.Stat.Name, stat.BaseStat)
		}
		formattedInfo = fmt.Sprintf("%sTypes:\n", formattedInfo)
		for _, pokeType := range info.Types {
			formattedInfo = fmt.Sprintf("%s  - %s\n", formattedInfo, pokeType.Type.Name)
		}
		fmt.Print(formattedInfo)
	}
	return nil
}

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

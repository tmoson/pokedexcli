package main

import (
	"fmt"
	"io"
	"strings"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type PokedexStyles struct {
	title        lipgloss.Style
	pokemon      lipgloss.Style
	selectedItem lipgloss.Style
	pagination   lipgloss.Style
	help         lipgloss.Style
	quitText     lipgloss.Style
}

func newStyles(darkBG bool) PokedexStyles {
	return PokedexStyles{
		title:        lipgloss.NewStyle().PaddingLeft(2).PaddingTop(5),
		pokemon:      lipgloss.NewStyle().PaddingLeft(4),
		selectedItem: lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170")),
		pagination:   list.DefaultStyles(darkBG).PaginationStyle.PaddingLeft(4),
		help:         list.DefaultStyles(darkBG).HelpStyle.PaddingLeft(4).PaddingBottom(1),
		quitText:     lipgloss.NewStyle().Margin(1, 0, 2, 4),
	}
}

type PokedexEntry struct {
	name   string
	number int
}

func (p PokedexEntry) FilterValue() string { return "" }

type pokemonDelegate struct {
	styles *PokedexStyles
}

func (d pokemonDelegate) Height() int                             { return 1 }
func (d pokemonDelegate) Spacing() int                            { return 0 }
func (d pokemonDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d pokemonDelegate) Render(w io.Writer, m list.Model, index int, pokedexEntry list.Item) {
	i, ok := pokedexEntry.(PokedexEntry)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", i.number, i.name)

	fn := d.styles.pokemon.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return d.styles.selectedItem.Render("> " + strings.Join(s, " "))
		}
	}
	fmt.Fprint(w, fn(str))
}

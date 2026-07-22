package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/tmoson/pokedexcli/internal/pokecache"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type ErrMsg struct{ err error }

func (e ErrMsg) Error() string { return e.err.Error() }

type UpdateMsg string

type Styles struct {
	BorderColor color.Color
	InputField  lipgloss.Style
}

func DefaultStyles() *Styles {
	s := new(Styles)
	s.BorderColor = lipgloss.Color("36")
	s.InputField = lipgloss.NewStyle().
		BorderForeground(s.BorderColor).
		BorderStyle(lipgloss.DoubleBorder()).
		Padding(1).Width(120)
	return s
}

type Model struct {
	commands      map[string]teaCommand
	configuration *config
	textInput     textinput.Model
	response      string
	width         int
	height        int
	styles        *Styles
}

func InitialModel() Model {
	ti := textinput.New()
	ti.Placeholder = "help, exit, map, explore, catch..."
	ti.SetWidth(120)
	ti.Prompt = "Pokedex > "
	ti.Focus()
	commands := getCommandsTea()
	configuration := config{
		nextLocation: "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		cache:        pokecache.NewCache(5 * time.Second),
		pokedex:      make(map[string]Pokemon),
	}
	styles := DefaultStyles()
	return Model{
		commands:      commands,
		configuration: &configuration,
		textInput:     ti,
		styles:        styles,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			command := cleanInput(m.textInput.Value())
			action, ok := m.commands[command[0]]
			if !ok {
				m.response = fmt.Sprintf("Unrecognized command: %v\n", command[0])
				m.textInput.Reset()
				return m, m.textInput.Focus()
			}
			m.response = action.callback(m.configuration, command[1:]...)
			m.textInput.Reset()
			return m, m.textInput.Focus()
		}
	}
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m Model) View() tea.View {
	content := lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Center,
			m.styles.InputField.Render(m.textInput.View()),
			m.response,
		),
	)
	v := tea.NewView(content)
	v.AltScreen = true
	return v
}

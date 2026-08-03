package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/tmoson/pokedexcli/internal/pokecache"

	"charm.land/bubbles/v2/list"
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
	commands       map[string]teaCommand
	configuration  *config
	textInput      textinput.Model
	response       string
	width          int
	height         int
	inputInd       int
	styles         *Styles
	pokedexStyles  PokedexStyles
	pokemonList    list.Model
	previousInputs []string
	pokedexList    []list.Item
}

func InitialModel() Model {
	pokedexStyles := newStyles(true)
	ti := textinput.New()
	ti.Placeholder = "help, exit, map, explore, catch..."
	ti.SetWidth(120)
	ti.Prompt = "PokeCLI > "
	ti.Focus()
	commands := getCommandsTea()
	configuration := config{
		nextLocation: "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		cache:        pokecache.NewCache(10 * time.Minute),
		pokedex:      make(map[string]Pokemon),
	}
	styles := DefaultStyles()
	return Model{
		commands:       commands,
		configuration:  &configuration,
		textInput:      ti,
		styles:         styles,
		pokedexStyles:  pokedexStyles,
		inputInd:       -1,
		previousInputs: []string{},
		pokedexList:    nil,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) UpdateList(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			m.pokedexList = nil
			m.textInput.Reset()
			m.textInput.Focus()
			return m, textinput.Blink
		case "enter":
			mon, ok := m.pokemonList.SelectedItem().(PokedexEntry)
			if ok {
				m.response = commandInspectTea(m.configuration, mon.name)
			}
			m.pokedexList = nil
			m.textInput.Reset()
			m.textInput.Focus()
			return m, textinput.Blink
		}
	}
	m.pokemonList, cmd = m.pokemonList.Update(msg)
	return m, cmd
}

func (m Model) GetPokedexEntries() []list.Item {
	var entries []list.Item
	for pokemon, pokemonInfo := range m.configuration.pokedex {
		newEntry := PokedexEntry{
			name:   pokemon,
			number: pokemonInfo.Id,
		}
		i := 0
		numEntries := len(entries)
		for ; i < numEntries; i++ {
			pokedexEntry := entries[i].(PokedexEntry)
			if pokedexEntry.number > pokemonInfo.Id {
				break
			}
		}
		// Still trying to understand why slices.Insert() won't work, but I think it's
		// because I have to return a slice of an interface, and that doesn't work well
		// with the generics implementation
		entries = append(entries[:i], append([]list.Item{newEntry}, entries[i:]...)...)
	}
	return entries
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		if m.pokedexList != nil {
			m.pokemonList.SetWidth(msg.Width)
		}
	case tea.KeyMsg:
		if m.pokedexList != nil {
			return m.UpdateList(msg)
		}
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			m.previousInputs = append(m.previousInputs, m.textInput.Value())
			m.inputInd = len(m.previousInputs) - 1
			command := cleanInput(m.textInput.Value())
			action, ok := m.commands[command[0]]
			if !ok {
				m.response = fmt.Sprintf("Unrecognized command: %v\n", command[0])
				m.textInput.Reset()
				return m, m.textInput.Focus()
			}
			// workaround currently, I think this should probably
			// be reworked to be handled as a new tea.Msg type
			switch action.name {
			case "pokedex":
				m.pokedexList = m.GetPokedexEntries()
				m.pokemonList = list.New(m.pokedexList, pokemonDelegate{}, m.width, m.height)
				m.pokemonList.Title = "Your Pokedex:"
				m.pokemonList.Styles.Title = m.pokedexStyles.title
				m.pokemonList.Styles.PaginationStyle = m.pokedexStyles.pagination
				m.pokemonList.Styles.HelpStyle = m.pokedexStyles.help
				m.pokemonList.SetShowStatusBar(false)
				m.pokemonList.SetFilteringEnabled(false)
				m.pokemonList.SetDelegate(pokemonDelegate{styles: &m.pokedexStyles})
				m.response = ""
				return m, nil
			case "exit":
				return m, tea.Quit
			default:
				m.response = action.callback(m.configuration, command[1:]...)
				m.textInput.Reset()
				return m, m.textInput.Focus()
			}
		case "up":
			if m.inputInd >= 0 {
				if m.inputInd > 0 {
					m.inputInd--
				}
				m.textInput.SetValue(m.previousInputs[m.inputInd])
			}
		case "down":
			if m.inputInd < len(m.previousInputs)-1 {
				m.inputInd++
				m.textInput.SetValue(m.previousInputs[m.inputInd])
			} else {
				m.textInput.Reset()
			}
		}
	}
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m Model) View() tea.View {
	var content string
	if m.pokedexList != nil {
		content = lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Center,
			lipgloss.Center,
			lipgloss.JoinVertical(
				lipgloss.Center,
				m.pokemonList.View(),
			),
		)
	} else {
		content = lipgloss.Place(
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
	}
	v := tea.NewView(content)
	v.AltScreen = true
	return v
}

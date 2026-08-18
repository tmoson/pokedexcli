package main

import (
	"encoding/json"
	"github.com/tmoson/pokedexcli/internal/pokecache"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestCommandInspectTea(t *testing.T) {
	mockPokedex := map[string]Pokemon{
		"pikachu": {
			Name:   "pikachu",
			Height: 4,
			Weight: 60,
			Stats: []PokemonStat{
				{
					BaseStat: 35,
					Stat:     NamedResource{Name: "hp"},
				},
				{
					BaseStat: 55,
					Stat:     NamedResource{Name: "attack"},
				},
			},
			Types: []PokemonType{
				{
					Type: NamedResource{Name: "electric"},
				},
			},
		},
	}

	conf := &config{
		Pokedex: mockPokedex,
	}

	cases := []struct {
		name     string
		inputs   []string
		expected []string
	}{
		{
			name:     "No inputs",
			inputs:   []string{},
			expected: []string{"Need a pokemon to inspect!"},
		},
		{
			name:     "Uncaught pokemon",
			inputs:   []string{"charmander"},
			expected: []string{"you have not caught charmander!\n"},
		},
		{
			name:   "Caught pokemon",
			inputs: []string{"pikachu"},
			expected: []string{
				"Name: pikachu",
				"Height: 0.40 M",
				"Weight: 0.60 kg",
				"Stats:",
				"  - hp: 35",
				"  - attack: 55",
				"Types:",
				"  - electric",
			},
		},
	}

	for _, c := range cases {
		actual := commandInspectTea(conf, c.inputs...)
		for _, expectedStr := range c.expected {
			if !strings.Contains(actual, expectedStr) {
				t.Errorf("Test %q failed: Expected output to contain %q, but got:\n%s", c.name, expectedStr, actual)
			}
		}
	}
}

func TestCommandHelpTea(t *testing.T) {
	conf := &config{}
	out := commandHelpTea(conf)

	expectedStrings := []string{
		"Welcome to the Pokedex!",
		"Usage:",
		"help",
		"catch",
		"inspect",
		"map",
		"mapb",
		"explore",
		"pokedex",
		"exit",
		"save",
		"load",
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(out, expected) {
			t.Errorf("Expected help output to contain %q", expected)
		}
	}
}

func TestCommandCatchTea(t *testing.T) {

	conf := &config{
		Cache:   pokecache.NewCache(5 * time.Minute),
		Pokedex: make(map[string]Pokemon),
	}
	// Create a pokemon guaranteed to be caught (BaseExperience = -1000)
	catchPokemon := Pokemon{
		Name:           "pidgey",
		BaseExperience: -1000,
	}
	catchBytes, _ := json.Marshal(catchPokemon)
	conf.Cache.Add("https://pokeapi.co/api/v2/pokemon/pidgey", catchBytes)

	// Create a pokemon guaranteed to escape (BaseExperience = 1000)
	escapePokemon := Pokemon{
		Name:           "mewtwo",
		BaseExperience: 1000,
	}
	escapeBytes, _ := json.Marshal(escapePokemon)
	conf.Cache.Add("https://pokeapi.co/api/v2/pokemon/mewtwo", escapeBytes)

	cases := []struct {
		name     string
		inputs   []string
		expected []string
	}{
		{
			name:     "No inputs",
			inputs:   []string{},
			expected: []string{"Can't catch nothing!"},
		},
		{
			name:     "Invalid ball",
			inputs:   []string{"pidgey", "master"},
			expected: []string{"Only have Pokeballs, Great Balls, and Ultra Balls"},
		},
		{
			name:     "Catch success",
			inputs:   []string{"pidgey"},
			expected: []string{"Throwing Pokeball at pidgey", "pidgey was caught!"},
		},
		{
			name:     "Catch success with Great Ball",
			inputs:   []string{"pidgey", "great"},
			expected: []string{"Throwing Great Ball at pidgey", "pidgey was caught!"},
		},
		{
			name:     "Catch escape",
			inputs:   []string{"mewtwo", "ultra"},
			expected: []string{"Throwing Ultra Ball at mewtwo", "mewtwo escaped!"},
		},
	}

	for _, c := range cases {
		out := commandCatchTea(conf, c.inputs...)
		for _, exp := range c.expected {
			if !strings.Contains(out, exp) {
				t.Errorf("Test %q failed: Expected output to contain %q, but got %q", c.name, exp, out)
			}
		}
	}
}

func TestCommandSaveLoadTea(t *testing.T) {
	tempDir := t.TempDir()
	t.Setenv("HOME", tempDir)

	conf1 := &config{
		NextLocation: "next_loc_123",
	}

	// Test save
	outSave := commandSaveTea(conf1)
	if !strings.Contains(outSave, "saved!") {
		t.Errorf("Expected save to succeed, got: %v", outSave)
	}

	// Verify file was created
	savePath := filepath.Join(tempDir, ".pokecliSaves", "save0")
	if _, err := os.Stat(savePath); os.IsNotExist(err) {
		t.Errorf("Save file was not created at expected path: %v", savePath)
	}

	// Test load
	conf2 := &config{}
	outLoad := commandLoadTea(conf2)
	if !strings.Contains(outLoad, "Loaded!") {
		t.Errorf("Expected load to succeed, got: %v", outLoad)
	}

	if conf2.NextLocation != conf1.NextLocation {
		t.Errorf("Expected loaded NextLocation to be %q, got %q", conf1.NextLocation, conf2.NextLocation)
	}

	// Test save with custom name
	outSaveCustom := commandSaveTea(conf1, "mysave")
	if !strings.Contains(outSaveCustom, "saved!") {
		t.Errorf("Expected custom save to succeed, got: %v", outSaveCustom)
	}

	// Test load with custom name
	conf3 := &config{}
	outLoadCustom := commandLoadTea(conf3, "mysave")
	if !strings.Contains(outLoadCustom, "Loaded!") {
		t.Errorf("Expected custom load to succeed, got: %v", outLoadCustom)
	}
	if conf3.NextLocation != conf1.NextLocation {
		t.Errorf("Expected loaded NextLocation to be %q, got %q", conf1.NextLocation, conf3.NextLocation)
	}
}

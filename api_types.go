package main

type NamedResource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Name struct {
	Name     string        `json:"name"`
	Language NamedResource `json:"language"`
}

type VersionDetails struct {
	Rate    int           `json:"rate"`
	Version NamedResource `json:"version"`
}

type EncounterMethodRate struct {
	EncounterMethod NamedResource    `json:"encounter_method"`
	VersionDetails  []VersionDetails `json:"version_details"`
}

type PokemonEncounter struct {
	Pokemon        NamedResource    `json:"pokemon"`
	VersionDetails []VersionDetails `json:"version_details"`
}

type LocationArea struct {
	Id                   int                   `json:"id"`
	Name                 string                `json:"name"`
	GameIndex            int                   `json:"game_index"`
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	Location             NamedResource         `json:"location"`
	Names                []Name                `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}

type LocationAPIResponse struct {
	Count    int             `json:"count"`
	Next     string          `json:"next"`
	Previous string          `json:"previous"`
	Results  []NamedResource `json:"results"`
}

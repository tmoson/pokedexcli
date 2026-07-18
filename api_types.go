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

type VersionGameIndex struct {
	GameIndex int           `json:"game_index"`
	Version   NamedResource `json:"version"`
}

type PokemonMoveVersion struct {
	MoveLearnMethod NamedResource `json:"move_learn_method"`
	VersionGroup    NamedResource `json:"version_group"`
	LevelLearnedAt  int           `json:"level_learned_at"`
	Order           int           `json:"order"`
}

type PokemonHeldItemVersion struct {
	Version NamedResource `json:"version"`
	Rarity  int           `json:"rarity"`
}

type EncounterMethodRate struct {
	EncounterMethod NamedResource    `json:"encounter_method"`
	VersionDetails  []VersionDetails `json:"version_details"`
}

type PokemonType struct {
	Slot int           `json:"slot"`
	Type NamedResource `json:"type"`
}

type PokemonStat struct {
	Stat     NamedResource `json:"stat"`
	Effort   int           `json:"effort"`
	BaseStat int           `json:"base_stat"`
}

type PokemonAbility struct {
	IsHidden bool          `json:"is_hidden"`
	Slot     int           `json:"slot"`
	Ability  NamedResource `json:"ability"`
}

type PokemonMove struct {
	Move                NamedResource        `json:"move"`
	VersionGroupDetails []PokemonMoveVersion `json:"version_group_details"`
}

type PokemonTypePast struct {
	Generation NamedResource `json:"generation"`
	Types      []PokemonType `json:"types"`
}

type PokemonAbilityPast struct {
	Generation NamedResource    `json:"generation"`
	Abilities  []PokemonAbility `json:"abilities"`
}

type PokemonStatPast struct {
	Generation NamedResource `json:"generation"`
	Stats      []PokemonStat `json:"stats"`
}

type PokemonSprites struct {
	FrontDefault     string `json:"front_default"`
	FrontShiny       string `json:"front_shiny"`
	FrontFemale      string `json:"front_female"`
	FrontShinyFemale string `json:"front_shiny_female"`
	BackDefault      string `json:"back_default"`
	BackShiny        string `json:"back_shiny"`
	BackFemale       string `json:"back_female"`
	BackShinyFemale  string `json:"back_shiny_female"`
}

type PokemonCries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}

type PokemonHeldItem struct {
	Item           NamedResource            `json:"item"`
	VersionDetails []PokemonHeldItemVersion `json:"version_details"`
}

type Pokemon struct {
	Id                     int                  `json:"id"`
	Name                   string               `json:"name"`
	BaseExperience         int                  `json:"base_experience"`
	Height                 int                  `json:"height"`
	IsDefault              bool                 `json:"is_default"`
	Order                  int                  `json:"order"`
	Weight                 int                  `json:"weight"`
	Abilities              []PokemonAbility     `json:"abilities"`
	GameIndices            []VersionGameIndex   `json:"game_indices"`
	HeldItem               PokemonHeldItem      `json:"held_item"`
	LocationAreaEncounters string               `json:"location_area_encounters"`
	Moves                  []PokemonMove        `json:"moves"`
	PastTypes              []PokemonTypePast    `json:"past_types"`
	PastAbilities          []PokemonAbilityPast `json:"pastAbilities"`
	PastStats              []PokemonStatPast    `json:"pastStats"`
	Sprites                PokemonSprites       `json:"sprites"`
	Cries                  []PokemonCries       `json:"cries"`
	Species                NamedResource        `json:"species"`
	Stats                  []PokemonStat        `json:"stats"`
	Types                  []PokemonType        `json:"types"`
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

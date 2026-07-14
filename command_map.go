package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type namedResource struct {
	Name string
	Url  string
}

type APIResponse struct {
	Count    int
	Next     string
	Previous string
	Results  []namedResource
}

func commandMap(conf *config) error {
	offset := conf.locationOffset
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%v&limit=20", offset)
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	var apiResponse APIResponse
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return err
	}
	for _, location := range apiResponse.Results {
		fmt.Printf("%s\n", location.Name)
	}
	conf.locationOffset += 20
	return nil
}

func commandMapb(conf *config) error {
	offset := conf.locationOffset
	if offset < 20 {
		return errors.New("Already at the beginning")
	}
	offset -= 20
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%v&limit=20", offset)
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	var apiResponse APIResponse
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return err
	}
	for _, location := range apiResponse.Results {
		fmt.Printf("%s\n", location.Name)
	}
	conf.locationOffset -= 20
	return nil
}

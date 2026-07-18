package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func commandMap(conf *config, inputs ...string) error {
	var offset int
	if len(inputs) == 0 {
		offset = conf.locationOffset
	} else {
		offset, parseErr := strconv.Atoi(inputs[0])
		if parseErr != nil {
			return parseErr
		}
		conf.locationOffset = offset
	}
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%v&limit=20", offset)
	var apiResponse LocationAPIResponse
	val, found := conf.cache.Get(url)
	if found {
		err := json.Unmarshal(val, &apiResponse)
		if err != nil {
			return err
		}
	} else {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		conf.cache.Add(url, body)
		err = json.Unmarshal(body, &apiResponse)
		if err != nil {
			return err
		}
	}
	for _, location := range apiResponse.Results {
		fmt.Printf("%s\n", location.Name)
	}
	conf.locationOffset += 20
	return nil
}

func commandMapb(conf *config, inputs ...string) error {
	var offset int
	if len(inputs) == 0 {
		offset = conf.locationOffset
	} else {
		offset, parseErr := strconv.Atoi(inputs[0])
		if parseErr != nil {
			return parseErr
		}
		conf.locationOffset = offset
	}
	if offset < 20 {
		return errors.New("Invalid back tracking index")
	}
	offset -= 20
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%v&limit=20", offset)
	var apiResponse LocationAPIResponse
	val, found := conf.cache.Get(url)
	if found {
		err := json.Unmarshal(val, &apiResponse)
		if err != nil {
			return err
		}
	} else {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		err = json.Unmarshal(body, &apiResponse)
		if err != nil {
			return err
		}
	}
	for _, location := range apiResponse.Results {
		fmt.Printf("%s\n", location.Name)
	}
	conf.locationOffset -= 20
	return nil
}

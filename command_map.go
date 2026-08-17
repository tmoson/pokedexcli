package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func commandMapTea(conf *config, inputs ...string) string {
	url := conf.NextLocation
	var apiResponse LocationAPIResponse
	val, found := conf.Cache.Get(url)
	if found {
		err := json.Unmarshal(val, &apiResponse)
		if err != nil {
			return err.Error()
		}
	} else {
		res, err := http.Get(url)
		if err != nil {
			return err.Error()
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err.Error()
		}
		conf.Cache.Add(url, body)
		err = json.Unmarshal(body, &apiResponse)
		if err != nil {
			return err.Error()
		}
	}
	output := ""
	for _, location := range apiResponse.Results {
		output = fmt.Sprintf("%s%s\n", output, location.Name)
	}
	conf.NextLocation = apiResponse.Next
	conf.PreviousLocation = apiResponse.Previous
	return output
}

func commandMapbTea(conf *config, inputs ...string) string {
	if conf.PreviousLocation == "" {
		return "Already at the beginning!"
	}
	url := conf.PreviousLocation
	var apiResponse LocationAPIResponse
	val, found := conf.Cache.Get(url)
	if found {
		err := json.Unmarshal(val, &apiResponse)
		if err != nil {
			return err.Error()
		}
	} else {
		res, err := http.Get(url)
		if err != nil {
			return err.Error()
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err.Error()
		}
		err = json.Unmarshal(body, &apiResponse)
		if err != nil {
			return err.Error()
		}
	}
	output := ""
	for _, location := range apiResponse.Results {
		output = fmt.Sprintf("%s%s\n", output, location.Name)
	}
	conf.NextLocation = apiResponse.Next
	conf.PreviousLocation = apiResponse.Previous
	return output
}

func commandMapb(conf *config, inputs ...string) error {
	if conf.PreviousLocation == "" {
		return errors.New("Already at the beginning!")
	}
	url := conf.PreviousLocation
	var apiResponse LocationAPIResponse
	val, found := conf.Cache.Get(url)
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
	conf.NextLocation = apiResponse.Next
	conf.PreviousLocation = apiResponse.Previous
	return nil
}

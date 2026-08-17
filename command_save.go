package main

import (
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"
)

func commandSaveTea(conf *config, inputs ...string) string {
	conf.lock.Lock()
	defer conf.lock.Unlock()
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Sprintf("Error finding home directory to save to: %v", err)
	}
	path := filepath.Join(homeDir, ".pokecliSaves")
	err = os.MkdirAll(path, 0755)
	if err != nil {
		return fmt.Sprintf("Error while creating save directory, make sure you have permission to save to ~/.pokecliSaves: %v", err)
	}
	if len(inputs) == 0 {
		path = filepath.Join(path, "save0")
	} else {
		path = filepath.Join(path, inputs[0])
	}
	file, err := os.Create(path)
	if err != nil {
		return fmt.Sprintf("Error while creating save file: %v", err)
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(conf)
	if err != nil {
		return fmt.Sprintf("Error while writing save to file: %v", err)
	}
	return "saved!"
}

func commandLoadTea(conf *config, inputs ...string) string {
	conf.lock.Lock()
	defer conf.lock.Unlock()
	path, err := os.UserHomeDir()
	if err != nil {
		return fmt.Sprintf("Error finding home directory to load save from: %v", err)
	}
	if len(inputs) == 0 {
		path = filepath.Join(path, ".pokecliSaves", "save0")
	} else {
		path = filepath.Join(path, ".pokecliSaves", inputs[0])
	}
	file, err := os.Open(path)
	if err != nil {
		return fmt.Sprintf("Error while opening save file: %v", err)
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(conf)
	if err != nil {
		return fmt.Sprintf("Error encountered loading save: %v", err)
	}
	return "Loaded!"
}

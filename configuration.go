package main

import (
	"encoding/json"
	"log"
	"os"
)

type configurationType struct {
	HtmlRoot string
	Listener string
}

var configuration = configurationType{}

func buildConfiguration() {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)

	err := decoder.Decode(&configuration)
	if err != nil {
		log.Fatal("Error loading configuration")
	}
}

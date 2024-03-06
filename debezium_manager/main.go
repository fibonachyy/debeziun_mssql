package main

import (
	"debezium_manager/domain"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

type ConfigData struct {
	Connectors []domain.Connector `json:"connectors"`
}

func main() {
	// Define command-line flag for the config file name
	configFile := flag.String("config", "config.json", "Path to the config JSON file")
	flag.Parse()

	// Read JSON file
	data, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	// Unmarshal JSON data into array of DebeziumConfig structs
	var configs []domain.Connector
	if err := json.Unmarshal(data, &configs); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	// Print each config
	for i, config := range configs {
		fmt.Printf("Config %d:\n", i+1)
		fmt.Printf("Name: %s\n", config.Name)
		fmt.Printf("Source: %s\n", config.Config.Source)
		fmt.Printf("Connector Class: %s\n", config.Config.ConnectorClass)
		// Add other fields as needed
		fmt.Println()
	}
}

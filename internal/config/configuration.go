package config

import (
	"encoding/json"
	"log"
	"os"
)

type Provider struct {
	OnlyAllowListed string         `json:"only_allow_listed"`
	List            []ProviderList `json:"list"`
}
type ProviderList struct {
	Name                   string     `json:"name"`
	Source                 string     `json:"source"`
	Version                string     `json:"version"`
	Allowed                bool       `json:"allowed"`
	ConfigurationParameter []Property `json:"configuration_parameters"`
}

type Property struct {
	Name  string  `json:"name"`
	Value *string `json:"value"`
}

type Configuration struct {
	Provider *Provider `json:"provider"`
}

func Parse(configFilePath string) (*Configuration, error) {
	// Read the JSON file
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	var config Configuration
	if err = json.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

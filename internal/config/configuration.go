package config

import (
	"encoding/json"
	"os"
)

// Provider holds terraform provider standardization rules/policies
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

// Variable holds policies used for standardization of variables
type Variable struct {
	Description Content `json:"description"`
	Name        Content `json:"name"`
}

type Content struct {
	Value string `json:"value"`
}

// Resource holds resource section configuration
type Resource struct {
	Type          string        `json:"type"`
	Attributes    []Property    `json:"attributes"`
	ChangeActions ChangeActions `json:"change_actions"`
	Allowed       Allowed       `json:"allowed"`
}

type ChangeActions struct {
	Update *CronExpression `json:"update"`
	Delete *CronExpression `json:"delete"`
	Create *CronExpression `json:"create"`
}
type CronExpression struct {
	Expression string `json:"time"`
}
type Allowed struct {
	When When `json:"when"`
}
type When struct {
	Expression string `json:"expression"`
}
type Configuration struct {
	Provider *Provider  `json:"provider"`
	Variable Variable   `json:"variables"`
	Resource []Resource `json:"resources"`
}

//TODO: review error handler

func Parse(configFilePath string) (*Configuration, error) {
	// Read the JSON file
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	var config Configuration
	if err = json.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type BodyguardsConfig struct {
	Run struct {
		Architecture string   `yaml:"architecture"`
		Enable       []string `yaml:"enable"`
	} `yaml:"run"`
}

func ReadBodyguardsYAML() (*BodyguardsConfig, error) {
	data, err := os.ReadFile("bodyguards.yaml")
	if err != nil {
		return nil, fmt.Errorf("failed to read YAML file: %v", err)
	}

	var config BodyguardsConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML: %v", err)
	}

	return &config, nil
}

func CreateBodyguardsYAML() error {
	config := BodyguardsConfig{
		Run: struct {
			Architecture string   `yaml:"architecture"`
			Enable       []string `yaml:"enable"`
		}{
			Architecture: "clean",
			Enable: []string{
				"linter",
				"report",
				"docs",
			},
		},
	}

	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to marshal YAML: %v", err)
	}

	err = os.WriteFile("bodyguards.yaml", data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write YAML file: %v", err)
	}

	return nil
}

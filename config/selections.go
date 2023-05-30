package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Run struct {
		Architecture string   `yaml:"architecture"`
		Enable       []string `yaml:"enable"`
	} `yaml:"run"`
}

func LoadSelections() (string, []string) {
	filename := "bodyguards.yaml"
	yamlContent, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(yamlContent, &config)
	if err != nil {
		log.Fatalf("Failed to unmarshal YAML: %v", err)
	}

	return config.Run.Architecture, config.Run.Enable
}

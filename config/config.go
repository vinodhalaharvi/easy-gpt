package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

// Config represents the top-level configuration structure.
type Config struct {
	Chat      ChatConfig      `yaml:"chat"`
	Assistant AssistantConfig `yaml:"assistant"`
}

type Message struct {
	Role    string `yaml:"role" json:"role"`
	Content string `yaml:"content" json:"content"`
}

// ChatConfig holds configuration specific to chat interactions.
type ChatConfig struct {
	RequestMethod string    `yaml:"requestMethod"`
	RequestURL    string    `yaml:"requestURL"`
	Model         string    `yaml:"model"`
	Messages      []Message `yaml:"messages"`
}

// AssistantConfig holds configuration specific to assistant interactions.
type AssistantConfig struct {
	SystemMessage string `yaml:"systemMessage"`
	RequestMethod string `yaml:"requestMethod"`
	RequestURL    string `yaml:"requestURL"`
	Model         string `yaml:"model"`
}

func (c *Config) FromYAMLFile(configPath string) *Config {
	// Default configuration path
	//configPath := "config.yaml"

	// Read the YAML configuration file
	data, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		os.Exit(1)
	}

	// Unmarshal the YAML into the Config struct
	config := &Config{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		fmt.Printf("Error parsing config file: %v\n", err)
		os.Exit(1)
	}

	return config
}

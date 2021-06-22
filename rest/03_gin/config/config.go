package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Mongo struct {
		Host     string `yaml:"host"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	} `yaml:"mongo"`
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
}

func GetConfigFromFile(configPath string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

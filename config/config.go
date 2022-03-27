package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	MySQL struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	} `yaml:"mysql"`
}

func GetConfig() (*Config, error) {
	config := &Config{}
	configPath := "./config.yml"

	s, err := os.Stat(configPath)

	if err != nil {
		return nil, err
	}

	if s.IsDir() {
		return nil, fmt.Errorf("'%s' is a directory", configPath)
	}

	file, err := os.Open(configPath)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

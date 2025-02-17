package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}

	Database struct {
		Host     string `yaml:"host"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
		Port     string `yaml:"port"`
		Type     string `yaml:"type"`
	}
}

func (c *Config) Load(path string) (*Config, error) {
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("yamlFile.Get err   #%v ", err)
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return &config, nil
}

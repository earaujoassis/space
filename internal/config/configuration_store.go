package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Service struct with data to retrieve a Vault secret
type Service struct {
	Space struct {
		ConfigurationStore struct {
			Addr  string `yaml:"addr"`
			Token string `yaml:"token"`
			Path  string `yaml:"path"`
		} `yaml:"configuration_store"`
	} `yaml:"space"`
}

// LoadService loads a Vault secret
func (s *Service) LoadService(filepath string) {
	yamlFile, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(yamlFile), &s)
	if err != nil {
		panic(err)
	}
}

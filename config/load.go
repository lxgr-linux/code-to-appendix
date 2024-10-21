package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

func Load(fileName string) (*Config, error) {
	var config = Config{}
	file, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

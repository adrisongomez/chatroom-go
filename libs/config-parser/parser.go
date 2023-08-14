package config_parser

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-yaml/yaml"
)

// Parse an yaml config file and return a Config struct.
func Parse(configPath string) (*Config, error) {
	absolutePath, error := filepath.Abs(configPath)
	if error != nil {
		return nil, error
	}
	fileContent, error := os.ReadFile(absolutePath)
	fmt.Println(string(fileContent))
	if error != nil {
		return nil, error
	}
	config := Config{}
	error = yaml.Unmarshal(fileContent, &config)
	if error != nil {
		return nil, error
	}
	return &config, nil
}

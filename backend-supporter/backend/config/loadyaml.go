package config

import (
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v3"
)

func LoadConfig() (*AppConfig, error) {
	_, filename, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(filename)
	yamlPath := filepath.Join(basePath, "config.yaml")

	data, err := os.ReadFile(yamlPath)
	if err != nil {
		return nil, err
	}

	var cfg AppConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

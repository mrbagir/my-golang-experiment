package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/ini.v1"
)

type TargetEndpoint struct{ Prefix, Local string }

type Config struct {
	TargetDomain    string
	LocalDomain     string
	TargetEndpoints []TargetEndpoint
}

func (c *Config) Print() {
	fmt.Println("[Config]")
	fmt.Println("TargetDomain:", c.TargetDomain)
	fmt.Println("LocalDomain:", c.LocalDomain)

	fmt.Println("\n[Endpoints]")
	for _, endpoint := range c.TargetEndpoints {
		fmt.Printf("  - Prefix: %s, Local: %s\n", endpoint.Prefix, endpoint.Local)
	}
	fmt.Println("\n[Logs]")
}

func LoadConfig() (*Config, error) {
	cfgIni, err := loadFileIni("")
	if err != nil {
		return nil, fmt.Errorf("failed to load ini file: %w", err)
	}

	sConfig := cfgIni.Section("Config")
	sEndpoints := cfgIni.Section("Endpoints")

	targetEndpoints := []TargetEndpoint{}
	for _, key := range sEndpoints.Keys() {
		targetEndpoints = append(targetEndpoints, TargetEndpoint{Prefix: key.Name(), Local: key.Value()})
	}

	config := &Config{
		TargetDomain:    sConfig.Key("target_domain").String(),
		LocalDomain:     sConfig.Key("local_domain").String(),
		TargetEndpoints: targetEndpoints,
	}

	if sConfig.Key("print_config").MustBool(false) {
		config.Print()
	}

	return config, nil
}

func loadFileIni(filename string) (*ini.File, error) {
	if filename == "" {
		filename = strings.TrimSuffix(os.Args[0], filepath.Ext(os.Args[0])) + ".ini"
	} else if strings.HasPrefix(filename, "~/") {
		filename = filepath.Join(filepath.Dir(os.Args[0]), filename[2:])
	}

	return ini.Load(filename)
}

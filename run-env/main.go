package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <env-name>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	envName := strings.ToLower(os.Args[1])

	switch envName {
	case "help", "usage", "?", "-h", "--help":
		fmt.Printf("Usage: %s <env-name>\n", filepath.Base(os.Args[0]))
		fmt.Println("Available environments: dev, prestage, stage, prerelease, release")
		os.Exit(1)
	case "dev", "prestage", "stage", "prerelease", "release":
	default:
		fmt.Printf("Unknown environment: %s\n", envName)
		os.Exit(1)
	}

	config, err := LoadAllConfig("", "", envName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ChangeFileEnv(config.Envs)
}

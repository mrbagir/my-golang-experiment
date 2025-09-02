package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <env-name> <command> [args...]\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	envName := strings.ToLower(os.Args[1])
	command := os.Args[2]
	args := os.Args[3:]

	switch envName {
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

	envVars := os.Environ()
	for key, value := range config.Envs {
		envVars = append(envVars, fmt.Sprintf("%s=%s", key, value))
	}

	cmd := exec.Command(command, args...)
	cmd.Env = envVars
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		os.Exit(1)
	}
}

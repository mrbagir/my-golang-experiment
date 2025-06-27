package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	env := "dev"
	if len(os.Args) > 1 {
		switch strings.ToLower(os.Args[1]) {
		case "usage", "help", "-h", "--help":
			fmt.Println("Usage: forwarder [dev|prestage|stage|prerelease|release]")
			fmt.Println("Example: forwarder dev")
			return
		case "dev", "prestage", "stage", "prerelease", "release":
			env = strings.ToLower(os.Args[1])
		}
	}

	defer fmt.Println("Program stopped.")

	config, err := loadConfig("", env)
	if err != nil {
		return
	}

	if config.PrintConfig {
		fmt.Println("\n[Logs]")
	}

	err = forwardWithError(config)
	if err == nil {
	} else if err.Error() == "Unauthorized" {
		fmt.Println("Error: Unauthorized. Please relogin to your OpenShift cluster.")
	} else {
		fmt.Printf("Error: %s\n", err)
	}
}

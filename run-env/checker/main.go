package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	maxLength := 0
	for _, env := range os.Environ() {
		splitEnv := strings.SplitN(env, "=", 2)
		if len(splitEnv) != 2 {
			continue
		}

		if len(splitEnv[0]) > maxLength {
			maxLength = len(splitEnv[0])
		}
	}

	for _, env := range os.Environ() {
		splitEnv := strings.SplitN(env, "=", 2)
		if len(splitEnv) != 2 {
			continue
		}

		fmt.Printf("%-*s = %s\n", maxLength, splitEnv[0], splitEnv[1])
	}
}

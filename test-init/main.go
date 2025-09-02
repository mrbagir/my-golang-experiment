package main

import (
	"fmt"

	_ "test-init/a"
	_ "test-init/b"
	_ "test-init/b/b1"
	// _ "test-init/b/b2"
)

func main() {
	fmt.Println("Hello, World!")
}

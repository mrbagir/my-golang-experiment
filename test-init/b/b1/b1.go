package b1

import (
	"fmt"
	_ "test-init/b/b2"
)

func init() {
	fmt.Println("Initializing package b1")
}

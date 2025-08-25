package main

import (
	"fmt"

	"github.com/google/uuid"
)

// func functionA() {
// 	pc, file, line, ok := runtime.Caller(1) // 1 artinya caller langsung
// 	if !ok {
// 		fmt.Println("Could not get caller info")
// 		return
// 	}
// 	callerFunc := runtime.FuncForPC(pc)
// 	fmt.Printf("Called by: %s\nIn file: %s:%d\n", callerFunc.Name(), file, line)
// }

// func functionB() {
// 	functionA()
// }

// func functionC() {
// 	functionA()
// }

// func main() {
// 	functionB()
// 	functionC()
// }

func functionA(callerID string) {
	fmt.Println("functionA called by ID:", callerID)
}

func functionB() {
	id := uuid.NewString()
	fmt.Println("functionB UUID:", id)
	functionA(id)
}

func functionC() {
	id := uuid.NewString()
	fmt.Println("functionC UUID:", id)
	functionA(id)
}

func main() {
	functionB()
	functionC()
}

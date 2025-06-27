package main

import (
	"fmt"
	"testing"
)

func TestLoad(t *testing.T) {
	cfgIni, err := loadFileIni("config.example.ini")
	if err != nil {
		t.Fatalf("Failed to load ini file: %v", err)
	}

	sEndpoints := cfgIni.Section("Endpoints")
	for _, key := range sEndpoints.Keys() {
		fmt.Println("Key:", key.Name(), "Value:", key.Value())
	}
}

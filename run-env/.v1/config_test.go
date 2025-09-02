package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestLoadFileIni(t *testing.T) {
	data, err := loadFileIni("runenv.ini")
	if err != nil {
		t.Fatalf("Failed to load ini file: %v", err)
	}
	fmt.Println(strings.Join(data.Section("Config").Key("allow_envs").Strings(","), "|"))
}

func TestLoadFileEnv(t *testing.T) {
	data, err := loadFileEnv(".env")
	if err != nil {
		t.Fatalf("Failed to load env file: %v", err)
	}
	fmt.Println("|" + data["lol"] + "|")
}

func TestSaveEmbed(t *testing.T) {
	createNewIniFile("")
}

package main

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/joho/godotenv"
	"gopkg.in/ini.v1"
)

type Config struct {
	Env        string
	Namespace  string
	PodName    string
	AllowEnvs  map[string]bool
	IgnoreEnvs map[string]bool
	Envs       map[string]string
}

func (c Config) Print() {
	fmt.Println("[Config]")
	fmt.Println("Env:", c.Env)
	fmt.Println("Namespace:", c.Namespace)
	fmt.Println("PodName:", c.PodName)
	if len(c.IgnoreEnvs) > 0 {
		fmt.Println("IgnoreEnvs:")
		for key := range c.IgnoreEnvs {
			fmt.Printf("  - %s\n", key)
		}
	}
	if len(c.AllowEnvs) > 0 {
		fmt.Println("AllowEnvs:")
		for key := range c.AllowEnvs {
			fmt.Printf("  - %s\n", key)
		}
	}
	if len(c.Envs) > 0 {
		fmt.Println("\n[Env]")
		for key, value := range c.Envs {
			fmt.Printf("  - %s: %s\n", key, value)
		}
	}
	fmt.Println()
}

func LoadAllConfig(envFilename, iniFilename, env string) (*Config, error) {
	env = strings.ToLower(env)

	// Load from .env
	cfgEnv, err := loadFileEnv(envFilename)
	if err != nil {
		return nil, fmt.Errorf("failed to load env file: %w", err)
	}

	// Load from [appname].ini
	cfgIni, err := loadFileIni(iniFilename)
	if err != nil {
		if os.IsNotExist(err) {
			createNewIniFile(iniFilename)
			return nil, fmt.Errorf("config file not found")
		} else {
			return nil, fmt.Errorf("failed to load ini file: %w", err)
		}
	}

	ocp := strings.Split(cfgEnv["ocp."+env], ",")
	if len(ocp) != 2 {
		return nil, fmt.Errorf("invalid ocp.%s format in env file, expected 'namespace,podname'", env)
	}

	config := &Config{
		Env:       env,
		Namespace: ocp[0],
		PodName:   ocp[1],
		Envs:      make(map[string]string),
	}

	sConfig := cfgIni.Section("Config")
	sCostum := cfgIni.Section("Custom")
	sCostumEnv := cfgIni.Section("Custom." + env)

	// Load from yaml on Kubernetes
	cfgYAML, err := loadYAML(ocp[0], ocp[1], sConfig.Key("always_import_yaml").MustBool(false))
	if err != nil {
		return nil, fmt.Errorf("failed to load YAML config: %w", err)
	}

	ignoreRegex := sConfig.Key("ignore_envs_regex").MustString("")
	allowRegex := sConfig.Key("allow_envs_regex").MustString("\\S+")

	if _, err = regexp.Compile(ignoreRegex); err != nil {
		return nil, fmt.Errorf("failed to compile ignore_envs_regex: %w", err)
	}
	if _, err = regexp.Compile(allowRegex); err != nil {
		return nil, fmt.Errorf("failed to compile allow_envs_regex: %w", err)
	}

	var ignoreEnvs = make(map[string]bool)
	for _, key := range sConfig.Key("ignore_envs").Strings(",") {
		if key == "" {
			continue
		}
		ignoreEnvs[strings.ToLower(key)] = true
	}
	config.IgnoreEnvs = ignoreEnvs

	var allowEnvs = make(map[string]bool)
	for _, key := range sConfig.Key("allow_envs").Strings(",") {
		if key == "" {
			continue
		}
		allowEnvs[strings.ToLower(key)] = true
	}
	config.AllowEnvs = allowEnvs

	importToMap(&config.Envs, cfgYAML, ignoreEnvs, allowEnvs, ignoreRegex, allowRegex)
	importToMap(&config.Envs, cfgEnv, ignoreEnvs, allowEnvs, ignoreRegex, allowRegex)
	importToMap(&config.Envs, sCostum.KeysHash(), map[string]bool{}, map[string]bool{}, "", "")
	importToMap(&config.Envs, sCostumEnv.KeysHash(), map[string]bool{}, map[string]bool{}, "", "")

	if sConfig.Key("print_config").MustBool(false) {
		config.Print()
	}

	return config, nil
}

func loadFileEnv(filename string) (map[string]string, error) {
	if filename == "" {
		filename = ".env"
	} else if strings.HasPrefix(filename, "~/") {
		filename = filepath.Join(filepath.Dir(os.Args[0]), filename[2:])
	}

	return godotenv.Read(filename)
}

func loadFileIni(filename string) (*ini.File, error) {
	if filename == "" {
		filename = strings.TrimSuffix(os.Args[0], filepath.Ext(os.Args[0])) + ".ini"
	} else if strings.HasPrefix(filename, "~/") {
		filename = filepath.Join(filepath.Dir(os.Args[0]), filename[2:])
	}

	return ini.Load(filename)
}

func saveFileIni(filename string, data map[string]string) error {
	if filename == "" {
		filename = strings.TrimSuffix(os.Args[0], filepath.Ext(os.Args[0])) + ".ini"
	} else if strings.HasPrefix(filename, "~/") {
		filename = filepath.Join(filepath.Dir(os.Args[0]), filename[2:])
	}

	if _, err := os.Stat(filepath.Dir(filename)); os.IsNotExist(err) {
		if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
			return fmt.Errorf("failed to create directory for ini file: %w", err)
		}
	}

	cfg := ini.Empty()
	section := cfg.Section("Custom")
	for key, value := range data {
		section.Key(key).SetValue(value)
	}

	return cfg.SaveTo(filename)
}

func importToMap(target *map[string]string, data map[string]string, ignore map[string]bool, allow map[string]bool, ignoreRegex, allowRegex string) {
	if *target == nil {
		*target = make(map[string]string)
	}

	ignoreReg := regexp.MustCompile(ignoreRegex)
	allowReg := regexp.MustCompile(allowRegex)

	isAllow := len(allow) > 0
	for key, value := range data {
		if ignore[strings.ToLower(key)] {
			continue
		}
		if isAllow && !allow[strings.ToLower(key)] {
			continue
		}
		if !ignoreReg.MatchString(key) || allowReg.MatchString(key) {
			continue
		}
		(*target)[key] = value
	}
}

//go:embed runenv.example.ini
var exampleFile []byte

func createNewIniFile(filename string) error {
	if filename == "" {
		filename = strings.TrimSuffix(os.Args[0], filepath.Ext(os.Args[0])) + ".ini"
	}

	return os.WriteFile(filename, []byte(exampleFile), 0644)
}

package main

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/anthhub/forwarder"
	"gopkg.in/ini.v1"
	"k8s.io/client-go/util/homedir"
)

const (
	defaultRemotePort = 9090
	defaultNamespace  = "default"
)

type Config struct {
	PrintConfig    bool
	RemotePort     int
	Namespace      string
	KubeConfigPath string
	ConfigFilename string
	Options        []*forwarder.Option
}

func (c Config) Print() {
	fmt.Println("[Config]")
	fmt.Printf("RemotePort: %d\n", c.RemotePort)
	fmt.Printf("Namespace: %s\n", c.Namespace)
	fmt.Printf("KubeConfigPath: %s\n", c.KubeConfigPath)
	fmt.Printf("ConfigFilename: %s\n", c.ConfigFilename)

	fmt.Println("\n[PortForward]")
	if len(c.Options) == 0 {
		fmt.Println("No target pods defined.")
		return
	}

	for i, v := range c.Options {
		fmt.Printf("%d. %s:%d\n", i+1, v.ServiceName, v.LocalPort)
	}
}

func loadConfig(filename string, env string) (*Config, error) {
	if filename == "" {
		filename = strings.TrimSuffix(os.Args[0], filepath.Ext(os.Args[0])) + ".ini"
	}

	cfg, err := ini.Load(filename)
	if err != nil {
		if os.IsNotExist(err) {
			createNewIniFile(filename)
			fmt.Printf("Config file '%s' not found. A new example file has been created. Please edit it and try again.\n", filename)
		}
		return nil, err
	}

	env = strings.ToLower(env)
	sConfig := cfg.Section("Config")
	sPortForward := cfg.Section("PortForward." + env)

	config := &Config{
		PrintConfig:    sConfig.Key("print_config").MustBool(false),
		RemotePort:     sConfig.Key("remotePort").RangeInt(defaultRemotePort, 1, 9999),
		KubeConfigPath: sConfig.Key("kubeConfigPath").MustString(os.Getenv("KUBECONFIG")),
		ConfigFilename: filename,
		Namespace:      sPortForward.Key("namespace").MustString(defaultNamespace),
		Options:        make([]*forwarder.Option, 0),
	}

	if strings.HasPrefix(config.KubeConfigPath, "~") {
		config.KubeConfigPath = filepath.Join(homedir.HomeDir(), config.KubeConfigPath[1:])
	}

	ports := make(map[int]bool)
	for _, key := range sPortForward.KeyStrings() {
		localPort := sPortForward.Key(key).RangeInt(0, 1, 9999)
		if ports[localPort] || key == "namespace" || localPort == 0 {
			continue
		}
		ports[localPort] = true

		config.Options = append(config.Options, &forwarder.Option{
			LocalPort:   localPort,
			RemotePort:  config.RemotePort,
			Namespace:   config.Namespace,
			ServiceName: key,
		})
	}

	if config.PrintConfig {
		config.Print()
	}

	return config, nil
}

//go:embed forwarder.example.ini
var exampleFile []byte

func createNewIniFile(filename string) {
	if filename == "" {
		filename = strings.TrimSuffix(os.Args[0], filepath.Ext(os.Args[0])) + ".ini"
	}

	os.WriteFile(filename, []byte(exampleFile), 0644)
}

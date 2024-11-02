package config

import (
	"fmt"
	"io"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type ConfigSchema struct {
	Mode string `yaml:"mode"`

	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`

	Auth struct {
		Salt string `yaml:"salt"`

		JWT struct {
			Secret string        `yaml:"secret"`
			Expiry time.Duration `yaml:"expiry"`
		} `yaml:"jwt"`
	}

	Postgres struct {
		URI string `yaml:"uri"`
	} `yaml:"postgres"`

	Redis struct {
		Addr     string `yaml:"addr"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	} `yaml:"redis"`

	Metrics struct {
		Enabled bool `yaml:"metrics"`
	} `yaml:"metrics"`
}

var Config *ConfigSchema

func LoadConfig(filePath string) {
	configFile, err := os.Open(filePath)
	var config ConfigSchema

	if err != nil {
		fmt.Printf("error opening config file: %v\n", err)
		os.Exit(1)
	}
	defer configFile.Close()

	configData, err := io.ReadAll(configFile)
	if err != nil {
		fmt.Printf("error reading config file: %v\n", err)
		os.Exit(1)
	}

	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		fmt.Printf("error unmarshalling config file: %v\n", err)
		os.Exit(1)
	}

	Config = &config
}

package config

import (
	"os"
	"strings"

	"github.com/go-yaml/yaml"
	"github.com/kelseyhightower/envconfig"
)

const configLocation = "./config.yaml"

type Config struct {
	Server struct {
		Host string `yaml:"host", envconfig:"SERVER_HOST"`
		Port string `yaml:"port", envconfig:"SERVER_PORT"`
	} `yaml:"server"`
}

func ReadConfig() (Config, error) {
	var config Config
	err := readConfig(&config)
	if err != nil {
		return Config{}, err
	}

	err = readEnv(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func readConfig(config *Config) error {
	configFile, err := os.Open(configLocation)
	if err != nil {
		return err
	}
	defer configFile.Close()

	decoder := yaml.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		return err
	}

	return nil
}

func readEnv(config *Config) error {
	return envconfig.Process("", config)
}

func FormatPort(port string) string {
	return strings.Join([]string{":", port}, "")
}

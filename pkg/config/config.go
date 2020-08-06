package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port string `yaml:"port", envconfig:"SERVER_PORT"`
	} `yaml:"server"`
	Database struct {
		Driver   string `yaml:"driver", envconfig:"DB_DRIVER"`
		Username string `yaml:"username", envconfig:"DB_USERNAME"`
		Password string `yaml:"password", envconfig:"DB_PASSWORD"`
		Host     string `yaml:"host", envconfig:"DB_HOST"`
		Port     string `yaml:"port", envconfig:"DB_PORT"`
		DbName   string `yaml:"dbname", envconfig:"DB_DBNAME"`
	} `yaml:"database"`
}

func LoadConfig(yamlFilePath string) (*Config, error) {
	var cfg Config

	err := readYamlFile(yamlFilePath, &cfg)
	if err != nil {
		return nil, err
	}

	err = readEnvironmentVar(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func readYamlFile(filepath string, cfg *Config) error {
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		return err
	}

	return nil
}

func readEnvironmentVar(cfg *Config) error {
	err := envconfig.Process("", cfg)
	if err != nil {
		return err
	}
	return nil
}

package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type ServerConfig struct {
	Host string `yaml:"host" envconfig:"SERVER_HOST"`
	Port string `yaml:"port" envconfig:"SERVER_PORT"`
}

type DatabaseConfig struct {
	Driver   string `yaml:"driver" envconfig:"DATABASE_DRIVER"`
	Username string `yaml:"username" envconfig:"DATABASE_USERNAME"`
	Password string `yaml:"password" envconfig:"DATABASE_PASSWORD"`
	Host     string `yaml:"host" envconfig:"DATABASE_HOST"`
	Port     string `yaml:"port" envconfig:"DATABASE_PORT"`
	DbName   string `yaml:"dbname" envconfig:"DATABASE_NAME"`
}

type ApiConfig struct {
	Endpoint string `yaml:"endpoint" envconfig:"API_ENDPOINT"`
	Token    string `yaml:"token" envconfig:"API_TOKEN"`
}

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	API      ApiConfig      `yaml:"api"`
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

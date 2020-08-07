package flags

import (
	"flag"
	"fmt"
	"os"
)

func ParseFlags() (string, error) {
	var configPath string
	flag.StringVar(&configPath, "config", "./config.yml", "Path to configuration file")
	flag.Parse()

	if err := isPathValid(configPath); err != nil {
		return "", err
	}

	return configPath, nil
}

func isPathValid(filepath string) error {
	s, err := os.Stat(filepath)
	if err != nil {
		return err
	}

	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory", filepath)
	}
	return nil
}

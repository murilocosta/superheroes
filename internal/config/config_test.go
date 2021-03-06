package config

import (
	"os"
	"testing"
)

const (
	expectedPort       = "7357"
	expectedDbDriver   = "test_driver"
	expectedDbUser     = "test_user"
	expectedDbPassword = "test_pass"
	expectedEndpoint   = "api_endpoint"
	expectedToken      = "api_token"
)

func TestReadYamlFile(t *testing.T) {
	configFilePath := "./testdata/config_test.yml"

	var cfg Config
	err := readYamlFile(configFilePath, &cfg)
	if err != nil {
		t.Errorf("Could not read from config file:\n%s", err)
	}

	testReadConfigObject(t, &cfg)
}

func TestReadEnvironmentVar(t *testing.T) {
	os.Setenv("SERVER_PORT", expectedPort)
	os.Setenv("DATABASE_DRIVER", expectedDbDriver)
	os.Setenv("DATABASE_USERNAME", expectedDbUser)
	os.Setenv("DATABASE_PASSWORD", expectedDbPassword)
	os.Setenv("API_ENDPOINT", expectedEndpoint)
	os.Setenv("API_TOKEN", expectedToken)

	var cfg Config
	readEnvironmentVar(&cfg)
	testReadConfigObject(t, &cfg)
}

func testReadConfigObject(t *testing.T, cfg *Config) {
	if cfg.Server.Port != expectedPort {
		t.Errorf("Env variable 'SERVER_PORT' is not as expected: %s", cfg.Server.Port)
	}

	if cfg.Database.Driver != expectedDbDriver {
		t.Errorf("Env variable 'DATABASE_DRIVER' is not as expected: %s", cfg.Database.Driver)
	}

	if cfg.Database.Username != expectedDbUser {
		t.Errorf("Env variable 'DATABASE_USERNAME' is not as expected: %s", cfg.Database.Username)
	}

	if cfg.Database.Password != expectedDbPassword {
		t.Errorf("Env variable 'DATABASE_PASSWORD' is not as expected: %s", cfg.Database.Password)
	}

	if cfg.API.Endpoint != expectedEndpoint {
		t.Errorf("Env variable 'API_ENDPOINT' is not as expected: %s", cfg.API.Endpoint)
	}

	if cfg.API.Token != expectedToken {
		t.Errorf("Env variable 'API_TOKEN' is not as expected: %s", cfg.API.Token)
	}
}

package config

import "testing"

func TestParseConnectionURL(t *testing.T) {
	expected := "host=localhost port=7357 user=test_u password=test_p dbname=test_db"
	cfg := &Config{
		Database: DatabaseConfig{
			Host:     "localhost",
			Port:     "7357",
			Username: "test_u",
			Password: "test_p",
			DbName:   "test_db",
		},
	}

	conn, err := ParseConnectionURL(cfg)
	if err != nil {
		t.Errorf("Error parsing database connection string:\n%s", err)
	}

	if conn != expected {
		t.Errorf("The parsed connection string is different from expected: %s", conn)
	}
}

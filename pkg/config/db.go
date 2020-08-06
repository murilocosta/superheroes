package config

import (
	"bytes"
	"html/template"
)

func ParseConnectionURL(cfg *Config) (string, error) {
	f := "host={{.Host}} port={{.Port}} user={{.User}} password={{.Password}} dbname={{.DbName}}"
	tmpl, err := template.New("database").Parse(f)
	if err != nil {
		return "", err
	}

	var conn bytes.Buffer
	err = tmpl.Execute(&conn, cfg.Database)
	if err != nil {
		return "", err
	}

	return conn.String(), nil
}

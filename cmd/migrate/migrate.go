package main

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/murilocosta/superheroes/internal/config"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.LoadConfig("./configs/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := config.ParseMigrationConnectionURL(cfg)
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.New("file://db/migrations", conn)
	if err != nil {
		log.Fatal(err)
	}

	if err = m.Up(); err != nil {
		log.Fatal(err)
	}
}

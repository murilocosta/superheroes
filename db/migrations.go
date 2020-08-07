package main

import (
	"database/sql"

	"github.com/murilocosta/superheroes/pkg/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.LoadConfig("./configs/config.yml")
	if err != nil {
		panic(err)
	}

	conn, err := config.ParseMigrationConnectionURL(cfg)
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://db/migrations", "postgres", driver)
	if err != nil {
		panic(err)
	}

	m.Steps(2)
}

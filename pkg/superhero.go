package main

import (
	"log"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/murilocosta/superheroes/pkg/config"
	"github.com/murilocosta/superheroes/pkg/flags"
)

func main() {
	cfgPath, err := flags.ParseFlags()
	handleError(err)

	cfg, err := config.LoadConfig(cfgPath)
	handleError(err)

	conn, err := config.ParseConnectionURL(cfg)
	handleError(err)

	db, err := gorm.Open(cfg.Database.Driver, conn)
	handleError(err)
	defer db.Close()

	srv := &http.Server{
		Addr:         cfg.Server.Host + ":" + cfg.Server.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

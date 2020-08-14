package main

import (
	"log"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/murilocosta/superheroes/internal/config"
	"github.com/murilocosta/superheroes/internal/flags"
	"github.com/murilocosta/superheroes/internal/superhero"
	"github.com/murilocosta/superheroes/pkg/logger"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	cfgPath, err := flags.ParseFlags()
	handleError(err)

	cfg, err := config.LoadConfig(cfgPath)
	handleError(err)

	lw, err := logger.NewLogRotation(cfg.Logging.Path, cfg.Logging.FileName)
	handleError(err)
	log.SetOutput(lw)

	conn, err := config.ParseConnectionURL(cfg)
	handleError(err)

	db, err := gorm.Open(cfg.Database.Driver, conn)
	defer db.Close()
	handleError(err)

	app := newApplicationHandler(cfg, db)
	addr := cfg.Server.Host + ":" + cfg.Server.Port
	srv := &http.Server{
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      app,
	}

	log.Println("Application running on " + addr)
	log.Fatal(srv.ListenAndServe())
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func newApplicationHandler(cfg *config.Config, db *gorm.DB) http.Handler {
	r := superhero.NewSuperRepository(db)
	api := superhero.NewSuperApi(cfg.API.Endpoint, cfg.API.Token)
	s := superhero.NewService(api, r)
	ctrl := superhero.NewCtrl(s)
	return superhero.NewRouter(ctrl)
}

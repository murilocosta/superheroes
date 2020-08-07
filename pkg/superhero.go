package main

import (
	"log"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/murilocosta/superheroes/pkg/config"
	"github.com/murilocosta/superheroes/pkg/flags"
	"github.com/murilocosta/superheroes/pkg/superhero"

	_ "github.com/jinzhu/gorm/dialects/postgres"
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

	app := setupApp(cfg, db)
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
		panic(err)
	}
}

func setupApp(cfg *config.Config, db *gorm.DB) http.Handler {
	r := superhero.NewSuperRepository(db)
	api := superhero.NewSuperHeroApi(cfg.API.Endpoint, cfg.API.Token)
	s := superhero.NewSuperHeroService(api, r)
	ctrl := superhero.NewSuperHeroCtrl(s)
	return superhero.NewHeroRouter(ctrl)
}

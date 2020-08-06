package main

import (
	"fmt"

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

	fmt.Println(cfg)

	conn, err := config.ParseConnectionURL(cfg)
	handleError(err)

	db, err := gorm.Open(cfg.Database.Driver, conn)
	handleError(err)
	defer db.Close()

}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

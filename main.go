package main

import (
	"log"

	"github.com/kohge2/upsdct-server/config"
	"github.com/kohge2/upsdct-server/database"
	"github.com/kohge2/upsdct-server/web/router"
)

func main() {
	if config.Env == nil {
		log.Fatal("config.Env is nil")
	}

	db, err := database.GetDB(config.Env.DbDsn)
	if err != nil {
		log.Fatal("DB ERROR", err)
	}
	txn, err := database.NewTransaction(config.Env.DbDsn)
	if err != nil {
		log.Fatal("Txn ERROR", err)
	}

	r := router.NewRouter(db, txn)
	r.Run(":8080")
}

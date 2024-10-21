package main

import (
	"log"

	"github.com/NoNamePL/webWallet/database"
	"github.com/NoNamePL/webWallet/iternal/config"
	"github.com/NoNamePL/webWallet/iternal/handlers/wallet"
	"github.com/gin-gonic/gin"
)

func main() {
	// init server
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// init config
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	// connect to DB
	db, err := database.ConnectDB(cfg)

	if err != nil {
		log.Fatal(err)
	}

	// init handlers
	wallet.RegisterRouter(router, db)

	// run server
	router.Run()
}

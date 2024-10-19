package main

import (
	"log"

	"github.com/NoNamePL/webWallet/database"
	"github.com/NoNamePL/webWallet/iternal/config"
	"github.com/NoNamePL/webWallet/iternal/handlers/wallet"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

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
	_ = db

	router.GET("/api/v1/wallet", wallet.GetBalance)

	router.Run()
}

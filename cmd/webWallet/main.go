package main

import (
	server "github.com/NoNamePL/webWallet"
	"log"
)

func main() {
	// TODO: init config

	// TODO: init logger

	// TODO: init storage
	srv := new(server.Server)
	if err := srv.Run(":8080"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

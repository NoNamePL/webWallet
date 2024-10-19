package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/NoNamePL/webWallet/iternal/config"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func ConnectDB(cfg *config.Config) (*Storage, error) {

	// Connected to DB
	connStr := fmt.Sprintf("host=db user=%s password=%s dbname=%s sslmode=disable",
		cfg.User, cfg.Password, cfg.DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Creating base tables in the database
	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS wallet(
			valletId UUID DEFAULT gen_random_uuid(),
    		amount BIGINT
		)
	`)

	if err != nil {
		return nil, errors.New("can't prepere query of wallet table")
	}

	_ , err = stmt.Exec()
	if err != nil {
		return nil, errors.New("can't create wallet table")
	}


	return &Storage{db: db}, nil

}

package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/NoNamePL/webWallet/iternal/config"

	_ "github.com/lib/pq"
)

func ConnectDB(cfg *config.Config) (*sql.DB, error) {

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
			valletId uuid UNIQUE NOT NULL DEFAULT gen_random_uuid(),
    		amount INT
		)
	`)

	if err != nil {
		return nil, errors.New("can't prepere query of wallet table")
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, errors.New("can't create wallet table")
	}

	// Create index
	stmt, err = db.Prepare(`
		CREATE UNIQUE INDEX IF NOT EXISTS valIdx on wallet (valletId)
	`)

	if err != nil {
		return nil, errors.New("can't prepere query of wallet table")
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, errors.New("can't create index")
	}

	return db, nil

}

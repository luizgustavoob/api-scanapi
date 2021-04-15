package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewConnection(postgresUrl string) (*sql.DB, error) {
	log.Default().Printf("POSTGRES URL: %s\n", postgresUrl)

	db, err := sql.Open("postgres", postgresUrl)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

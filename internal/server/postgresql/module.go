package postgresql

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	New,
)

func New() (*sql.DB, error) {
	databaseURL := os.Getenv("POSTGRES_URL")

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

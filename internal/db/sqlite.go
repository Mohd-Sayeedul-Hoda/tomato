package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func OpenSqliteConnection() (*sql.DB, error) {

	db, err := sql.Open("sqlite", "../../tomato.db")
	if err != nil {
		return nil, err
	}

	return db, nil
}

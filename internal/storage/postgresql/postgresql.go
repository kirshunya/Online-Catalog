package postgresql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" //init postgres driver
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const fun = "storage.postgres.New"

	db, err := sql.Open("postgres", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s, %w", fun, err)
	}

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXIST url (
	    id INTEGER PRIMARY KEY 
	    alias TEXT NOT NULL UNIQUE
	    url TEXT NOT NULL);
	CREATE INDEX IF NOT EXISTS idx_alias ON url(alias);
	`)
	if err != nil {
		return nil, fmt.Errorf("%s, %w", fun, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s, %w", fun, err)
	}

	return &Storage{db: db}, nil
}

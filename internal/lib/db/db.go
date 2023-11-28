package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase(host, port, dbname, user, password string) (*Database, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, dbname, user, password)
	db, cErr := sql.Open("postgres", connStr)
	if cErr != nil {
		return nil, cErr
	}
	if pErr := db.Ping(); pErr != nil {
		return nil, pErr
	}
	return &Database{db: db}, nil
}

func (d *Database) Close() error {
	return d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}

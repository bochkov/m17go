package db

import (
	"database/sql"
	"fmt"
)

type Database struct {
	host     string
	port     string
	dbname   string
	user     string
	password string
}

func NewDatabase(host, port, dbname, user, password string) *Database {
	return &Database{host: host, port: port, dbname: dbname, user: user, password: password}
}

func (db Database) Connect() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		db.host, db.port, db.dbname, db.user, db.password)
	conn, cErr := sql.Open("postgres", connStr)
	if cErr != nil {
		return nil, cErr
	}
	if pErr := conn.Ping(); pErr != nil {
		return nil, pErr
	}
	return conn, nil
}

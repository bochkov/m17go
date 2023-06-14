package db

import (
	"database/sql"
	"fmt"
)

type database struct {
	host     string
	port     string
	dbname   string
	user     string
	password string
}

func NewDatabase(host, port, dbname, user, password string) *database {
	return &database{host: host, port: port, dbname: dbname, user: user, password: password}
}

func (db database) Connect() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		db.host, db.port, db.dbname, db.user, db.password)
	conn, cErr := sql.Open("postgres", connStr)
	if cErr != nil {
		return nil, cErr
	}
	pErr := conn.Ping()
	if pErr != nil {
		return nil, pErr
	}
	return conn, nil
}

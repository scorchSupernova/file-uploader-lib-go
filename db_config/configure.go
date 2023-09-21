package db_config

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func Connect(db_dsn string) (*sql.DB, error) {
	conn, err := sql.Open("postgres", db_dsn)
	if err != nil {
		return nil, err
	}
	log.Println("Db Connection established")
	return conn, nil
}

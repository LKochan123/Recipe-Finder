package db

import (
	"database/sql"
	"log"
	"path/filepath"
)

var DB *sql.DB 

func InitDB() {
	var err error
	path := filepath.Join("db", "test.db")

	DB, err = sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}
}
package database

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func GetDB() *sql.DB {
	db, err := sql.Open("sqlite", "./database.db")
	if err != nil {
		panic(err)
	}
	return db
}

func CloseDB(db *sql.DB) {
	db.Close()
}

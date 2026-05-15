package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("erro ao conectar no banco:", err)
	}

	db.Exec("PRAGMA journal_mode = WAL;")

	return db
}

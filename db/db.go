package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "api.sql")
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS EVENTS (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	location TEXT NOT NULL,
	dateTime DATETIME NOT NULL,
	user_id INTEGER
	)
	`
	_, err := DB.Exec(createEventsTable)
	if err != nil {
		log.Fatalf("Could not create events table: %v", err)
	}
}

func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Fatalf("Could not close the database: %v", err)
	}
}

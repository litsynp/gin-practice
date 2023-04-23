package db

import (
	"database/sql"
	"gin-practice/internal"

	_ "github.com/lib/pq"
)

type Database struct {
	*sql.DB
}

var db *Database

func GetDb() *Database {
	if db != nil {
		return db
	}

	database, err := sql.Open("postgres", internal.DB_URL)
	if err != nil {
		panic(err)
	}

	db = &Database{database}
	db.Migrate()

	return db
}

func (db *Database) Migrate() {
	if err := migrateUser(db); err != nil {
		panic(err)
	}
}

func CloseDb() {
	defer func() { db = nil }()

	if err := db.Close(); err != nil {
		panic(err)
	}
}

func migrateUser(db *Database) error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
	    id SERIAL PRIMARY KEY,
	    first_name VARCHAR(20),
	    last_name VARCHAR(20),
	    email TEXT NOT NULL UNIQUE,
	    password VARCHAR(255)
 	);
	`)

	return err
}

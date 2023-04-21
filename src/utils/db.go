package utils

import (
	"database/sql"
	"gin-practice/src/users"
	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func InitDb() *DB {
	db, err := sql.Open("postgres", DB_URL)
	checkError(err)

	err = db.Ping()
	checkError(err)

	database := &DB{db}
	database.Migrate()

	return database
}

func (db *DB) Migrate() {
	err := users.UserRepository.Migrate(db.DB)
	if err != nil {
		panic(err)
	}

	checkError(err)
}

func (db *DB) CloseDb() {
	err := db.Close()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

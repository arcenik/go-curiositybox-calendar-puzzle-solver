package dbsqlite3

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type dbStatement struct {
	id    string
	query string
}

type DB struct {
	db *sql.DB
}

func InitDB(fileName string) (db DB) {

	statements := []dbStatement{
		// {
		// 	"enable foreign keys",
		// 	"PRAGMA foreign_keys = ON;",
		// },
		{
			"solutions table",
			`CREATE TABLE IF NOT EXISTS solutions (
				id      INTEGER PRIMARY KEY AUTOINCREMENT,
				month   INTEGER,
				day     INTEGER,
				json    TEXT NOT NULL
			);`,
		},
		{
			"solutions date index",
			"CREATE INDEX IF NOT EXISTS solutions_date ON solutions(month, day);",
		},
		{
			"solutions json unique index",
			"CREATE UNIQUE INDEX IF NOT EXISTS solutions_json ON solutions(json);",
		},
	}

	if fileName == "" {
		panic("dbsqlite3.InitDB(\"\") called without filename")
	}
	var err error
	db.db, err = sql.Open("sqlite3", fileName)
	if err != nil {
		panic(err)
	}
	db.db.SetMaxOpenConns(10)
	db.db.SetMaxIdleConns(5)

	for _, statement := range statements {
		// fmt.Println("Apply SQL statement", statement.id)
		_, err := db.db.Exec(statement.query)
		if err != nil {
			panic(fmt.Sprintf("SQL statement %v failed: %v", statement.id, err))
		}
	}
	return db
}

func (db *DB) Close() error {
	return db.db.Close()
}

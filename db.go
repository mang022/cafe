package main

import (
	"database/sql"
	"time"
)

var db *sql.DB

func setupDB() {
	var err error
	db, err = sql.Open("mysql", "root:pwd@tcp(127.0.0.1:3306)/cafe")
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Duration(2 * time.Minute))
	db.SetMaxIdleConns(100)
}

func closeDB() {
	_ = db.Close()
}

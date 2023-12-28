package main

import (
	"database/sql"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func setupDB() {
	var err error
	db, err = sql.Open("mysql", buildDataSourceName())
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Duration(2 * time.Minute))
	db.SetMaxIdleConns(100)
}

func closeDB() {
	_ = db.Close()
}

func buildDataSourceName() string {
	return conf.DB.User + ":" + conf.DB.Pwd + "@tcp(" + conf.DB.Host + ":" + strconv.Itoa(conf.DB.Port) + ")/cafe"
}

package db

import (
	"database/sql"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/mang022/cafe/conf"
)

var CafeDB *sql.DB

func SetupDB() {
	var err error
	CafeDB, err = sql.Open("mysql", buildDataSourceName())
	if err != nil {
		panic(err)
	}

	CafeDB.SetConnMaxLifetime(time.Duration(2 * time.Minute))
	CafeDB.SetMaxIdleConns(100)
}

func CloseDB() {
	_ = CafeDB.Close()
}

func buildDataSourceName() string {
	return conf.Conf.DB.User + ":" +
		conf.Conf.DB.Pwd + "@tcp(" +
		conf.Conf.DB.Host + ":" +
		strconv.Itoa(conf.Conf.DB.Port) + ")/cafe"
}

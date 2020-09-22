package db

import (
	"database/sql"
	"time"

	"github.com/aak1247/AssertTiDB/config"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func GetDb() *sql.DB {
	if db == nil {
		env := config.GetEnv()
		var err error
		db, err = sql.Open("mysql", env.Tidb.GetUrl())
		if err != nil {
			panic(err)
		}
		// See "Important settings" section.
		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)
	}
	return db
}

func Close() error {
	return db.Close()
}

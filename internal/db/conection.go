package db

import (
	"database/sql"
	"time"
)

func NewDbConnection() *sql.DB {
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}

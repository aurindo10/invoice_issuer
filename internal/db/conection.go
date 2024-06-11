package db

import (
	"database/sql"
	"time"

	"github.com/aurindo10/invoice_issuer/pkg/utils"
	_ "github.com/go-sql-driver/mysql"
)

func NewDbConnection() *sql.DB {
	connString := utils.GetEnv("GOOSE_DBSTRING", "")
	db, err := sql.Open("mysql", connString)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}

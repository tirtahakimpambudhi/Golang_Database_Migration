package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go_database_migration/config"
	"log"
	"time"
)

func NewDB() *sql.DB {

	db, err := sql.Open(config.DIALECT, config.CONNECT)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	db.SetConnMaxLifetime(time.Minute * 60)
	db.SetConnMaxIdleTime(time.Minute * 5)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	return db
}

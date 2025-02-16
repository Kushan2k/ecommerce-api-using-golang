package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func NewMySqlDatabase(cfg mysql.Config) (*sql.DB, error) {
	
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
		
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

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

	er:=db.Ping()
	if er!=nil{
		log.Fatal("Error connecting to database: ", er)
	}
	log.Println("Connected to database")

	return db, nil
}

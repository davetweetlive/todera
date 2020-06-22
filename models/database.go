package models

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToMySQL() (*sql.DB, error) {

	db, err := sql.Open("mysql", "root:Megamind@1@(127.0.0.1:3306)/note?parseTime=true")

	if err != nil {
		log.Fatal("Couldn't connect to the database!")
	}

	return db, err
}

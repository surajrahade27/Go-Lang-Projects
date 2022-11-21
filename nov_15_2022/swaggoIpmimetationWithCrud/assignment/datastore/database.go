package datastore

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func DBCONN() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:Password@1@tcp(localhost:3306)/DEMO")

	if err != nil {
		log.Fatal(err)
	}
	return db

}

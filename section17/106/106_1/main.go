package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DbConnection *sql.DB

func main() {
	var err error
	DbConnection, err = sql.Open("sqlite3", "../example.sql")
	if err != nil {
		log.Fatalln(err)
	}
	defer DbConnection.Close()

	cmd := `CREATE TABLE IF NOT EXISTS persons(
		name TEXT,
		age  INTEGER
	)`
	_, err = DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("table ok")
}

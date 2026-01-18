package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DbConnection *sql.DB

func main() {
	// ① DBを開く（なければ作成される）
	var err error
	DbConnection, err = sql.Open("sqlite3", "../example.sql")
	if err != nil {
		log.Fatalln(err)
	}
	defer DbConnection.Close()


	// ③ INSERT
	cmd := "INSERT INTO persons (name, age) VALUES (?, ?)"
	_, err = DbConnection.Exec(cmd, "Nancy", 20)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("insert ok")
}

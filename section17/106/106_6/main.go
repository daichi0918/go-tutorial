package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DbConnection *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	// ① DBを開く
	var err error
	DbConnection, err = sql.Open("sqlite3", "../example.sql")
	if err != nil {
		log.Fatalln(err)
	}
	defer DbConnection.Close()

	// ② DELETE
	cmd := "DELETE FROM persons WHERE name = ?"
	result, err := DbConnection.Exec(cmd, "Nancy")
	if err != nil {
		log.Fatalln(err)
	}

	// ③ 影響行数を確認（おすすめ）
	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("deleted rows: %d\n", rows)
}

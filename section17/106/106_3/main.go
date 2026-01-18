package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DbConnection *sql.DB

func main() {
	// ① DBを開く
	var err error
	DbConnection, err = sql.Open("sqlite3", "../example.sql")
	if err != nil {
		log.Fatalln(err)
	}
	defer DbConnection.Close()

	// ③ UPDATE
	cmd := "UPDATE persons SET age = ? WHERE name = ?"
	result, err := DbConnection.Exec(cmd, 25, "Nancy")
	if err != nil {
		log.Fatalln(err)
	}

	// ④ 影響行数を確認（おすすめ）
	rows, _ := result.RowsAffected()
	log.Printf("updated rows: %d\n", rows)
}

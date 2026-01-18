package main

import (
	"database/sql"
	"fmt"
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

	// ② QueryRow（1件取得）
	cmd := "SELECT name, age FROM persons WHERE age = ?"
	row := DbConnection.QueryRow(cmd, 1000)

	// ③ Scan
	var p Person
	err = row.Scan(&p.Name, &p.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No row")
			return
		}
		log.Fatalln(err)
	}

	// ④ 結果表示
	fmt.Println(p.Name, p.Age)
}

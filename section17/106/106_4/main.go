package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// コネクションプール（学習用にグローバル）
var DbConnection *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	// ① DBを開く（前と同じ書き方）
	var err error
	DbConnection, err = sql.Open("sqlite3", "../example.sql")
	if err != nil {
		log.Fatalln(err)
	}
	defer DbConnection.Close()

	// ② SELECT
	cmd := "SELECT name, age FROM persons"
	rows, err := DbConnection.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	// ③ rows → struct に詰める
	var pp []Person
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Fatalln(err)
		}
		pp = append(pp, p)
	}

	// ④ rows のエラーチェック
	if err = rows.Err(); err != nil {
		log.Fatalln(err)
	}

	// ⑤ 結果表示
	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}
}

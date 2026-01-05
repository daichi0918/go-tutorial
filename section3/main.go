package main

import "fmt"

// g := 5

func main() {
	// 変数定義
	// var 変数名 型 = 値
	var i int = 100;
	fmt.Println(i);

	var s string = "Hello Go";
	fmt.Println(s);

	var t, f bool = true, false;
	fmt.Println(t, f);

	var (
		i2 int = 200
		s2 string = "Go Lang"
	)
	fmt.Println(i2, s2);

	// 暗黙的な定義
	// 変数名 := 値
	// 基本的には型指定した方がいい
	n := 300
	fmt.Println(n)

	// n = "Hello";
	// fmt.Println(n) ← コンパイルエラーになる

	fmt.Println(g)
}
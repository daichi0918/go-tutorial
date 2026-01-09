package main

import "fmt"

// 定数

const Pi = 3.14 // 最初を大文字にすると、パッケージ外からも参照可能

const (
	URL			= "https://example.com"
	SiteName	= "Example"
)

const (
	A = 1
	B
	C = "A"
	D
	E
)

const (
	c0 = iota
	c1
	c2
)

// var Big int = 929229292922222222 + 1
const Big = 929229292922222222 + 1

func main() {
	fmt.Println(Pi)

	// Pi = 3;
	// fmt.Println(Pi)

	fmt.Println(URL, SiteName)

	fmt.Println(A, B, C, D, E)

	fmt.Println(Big - 1)

	fmt.Println(c0, c1, c2)
}
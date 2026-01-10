package main

import "fmt"

func Plus(x int, y int) int {
	// func Plus(x, y int) int { // 同じ時はこう書ける
	return x + y
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func Double(price int) (result int) {
	result = price * 2
	return // return result と同じ(引数で指定している場合)
}

func noreturn() {
	fmt.Println("No Return")
	return
}


func main() {
	i:= Plus(3, 5)
	fmt.Println(i)

	i2, i3 := Div(10, 3)
	// i2, _ := Div(10, 3) // 2つ目の戻り値を無視する場合
	fmt.Println(i2, i3)

	i4 := Double(100)
	fmt.Println(i4)

	noreturn()
}
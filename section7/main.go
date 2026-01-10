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


func returnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

func callFunction(f func()) {
	f()
}

// クロージャー(外側の変数を保持し続ける関数)
func later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}


func main() {
	// 関数
	i:= Plus(3, 5)
	fmt.Println(i)

	i2, i3 := Div(10, 3)
	// i2, _ := Div(10, 3) // 2つ目の戻り値を無視する場合
	fmt.Println(i2, i3)

	i4 := Double(100)
	fmt.Println(i4)

	noreturn()

	// 無名関数
	f := func(x, y int) int {
		return x + y
	}
	i5 := f(10, 20)
	fmt.Println(i5)

	// 即時実行関数
	i6 :=	func(x, y int) int {
		return x + y
	}(1, 2)

	fmt.Println(i6)

	// 関数を返す関数
	f2:= returnFunc()
	f2()

	// 関数を引数に取る関数
	callFunction(func() {
		fmt.Println("I'm a function!!")
	}) // => "I'm a function"

	// クロージャーの実装
	fc := later()
	fmt.Println(fc("Hello"))
	fmt.Println(fc("my"))
	fmt.Println(fc("name"))
	fmt.Println(fc("is"))
	fmt.Println(fc("Gopher"))

	// ジェネレーターの実装
	ints := integers()

	fmt.Println(ints()) // => "1"
	fmt.Println(ints()) // => "2"
	fmt.Println(ints()) // => "3"

	otherInts := integers()
	fmt.Println(otherInts()) // => "1"
}
package main // goは一つのパッケージしか持てない

import (
	"fmt" 
	"strconv"
	) 


func main() {

	// 条件分岐
	a := 1
	//if  else if  else
	if a == 2 {
		fmt.Println("2だ")
	} else if a == 1 {
		fmt.Println("1だ")
	} else {
		fmt.Println("それ以外")
	}

		if b := 100; b == 100 {
		fmt.Println("100だ")
	}

	x := 5
	if x := 2; true {
		fmt.Println(x)
		//2
	}
	fmt.Println(x)

	// エラーハンドリング
		var s string = "100"
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T %v\n", i, i)
}
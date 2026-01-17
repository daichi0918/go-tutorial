package main

import (
	"fmt"

	"example.com/golang-webgosql/section14/go_test/alib"
)

func IsOne(i int) bool {
	if i == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	// fmt.Println(IsOne(1))
	// fmt.Println(IsOne(0))
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)
	fmt.Println(alib.Average(s))
}

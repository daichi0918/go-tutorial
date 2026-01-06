package main // goは一つのパッケージしか持てない

import (
	"fmt"  // formatの略
	"time" 
	) 


func main() {
	fmt.Println("Hello, World!")
	fmt.Println(time.Now())
}
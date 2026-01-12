package main

import "fmt"

func main() {

	// var ch1 chan int
// // var ch2 <-chan int  // 受信専用チャネル（receive-only）
// // var ch3 chan<- int  // 送信専用チャネル（send-only）


	// fmt.Println(cap(ch1))

	// ch1 = make(chan int, 20)
	// fmt.Println(cap(ch1))

	c1 := make(chan int)
	c2 := make(chan int, 10)
	fmt.Println(cap(c1))

	c2 <- 1
	fmt.Println(len(c2))
	i := <-c2
	fmt.Println(i)
	fmt.Println(len(c2))

	c2 <- 2
	c2 <- 3
	c2 <- 4
	fmt.Println(len(c2))

	fmt.Println(<-c2)
	fmt.Println(len(c2))

	fmt.Println(<-c2)
	fmt.Println(len(c2))

	fmt.Println(<-c2)
	fmt.Println(len(c2))

}

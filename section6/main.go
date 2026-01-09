package main

import "fmt"



func main() {
	/* 算術演算子 */
	fmt.Println(1+2);
	fmt.Println("ABC" + "DEF");

	fmt.Println(5-3);
	fmt.Println(4*3);
	fmt.Println(10/2);
	fmt.Println(7%3);

	n:= 0;
	n +=4; // n= n + 4 と同じ
	fmt.Println(n)
	n++ // n = n + 1 と同じ
	fmt.Println(n)
	n--
	fmt.Println(n)

	s:= "Hello ";
	s += "Go";
	fmt.Println(s);

	/* 比較演算子 */
	fmt.Println(1 == 1);
	fmt.Println(1 == 2);
	fmt.Println(1 <= 2);
	fmt.Println(1 > 2);

	fmt.Println(true == false);
	fmt.Println(true != false)

	/* 論理演算子 */
	fmt.Println(true && false == true);
	fmt.Println(true && true == true);
	fmt.Println(true || false == true);
	fmt.Println(false || false == true);
	fmt.Println(!true);
}

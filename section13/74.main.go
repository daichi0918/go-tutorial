package main

import (
	"fmt"

	"example.com/golang-webgosql/section13/foo"
)

func appName() string {
	const AppName = "Go App"
	var Version = "1.0"
	return AppName + " " + Version
}

func Do(s string) (b string) {
	b = s
	{
		b := "BBBB"
		fmt.Println(b) // ブロック内のbが優先される	
	}
	fmt.Println(b)
	return b
}

func main() {
	fmt.Println(foo.MAX);
	// fmt.Println(foo.min); 

	fmt.Println(foo.ReturnMin());

	fmt.Println(appName())

	// fmt.Println(AppName) // 呼び出せない

	fmt.Println(Do("AAAA"))
}

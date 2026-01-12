package main

import "fmt"

func main() {
	sl := []string{"A", "B", "C"}
	fmt.Println(sl)


	for i, v := range sl {
			fmt.Println(i, v)
	}


	for _, v := range sl {
			fmt.Println(v)
	}

	for i := range sl {
			fmt.Println(i)
	}


	for i := 0; i < len(sl); i++ {
		fmt.Println(i, sl[i])
	}

	// lenとrangeの違い
	// lenは毎回関数を呼び出すが、rangeは最初に長さを取得している
	// →つまり、rangeの方がパフォーマンスが良い
	
	// 無限ループにならない
// 	for i := range sl {
//     sl = append(sl, "X")
// }

// 無限ループになる
// for i := 0; i < len(sl); i++ {
//     sl = append(sl, "X")
// }


}

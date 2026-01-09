package main

import "fmt"
import "strconv"

func main() {
/* 	整数型 */
	fmt.Println("整数型");
	var i int = 100;

	fmt.Println(i+50);

	// 型を調べる
	fmt.Printf("%T\n", i);

	// 型変換
	fmt.Printf("%T\n", int32(i));

	/* 浮動小数点型 */
	fmt.Println("浮動小数点型");

	var f float64 = 2.4;
	fmt.Println(f);

	fl := 3.2;
	fmt.Println(f + fl);
	fmt.Printf("%T, %T\n", fl, f);

	// float32は基本的に使わない
	var fl32 float32 = 1.2;
	fmt.Printf("%T\n", fl32);

	fmt.Printf("%T\n", float64(fl32));

	zero := 0.0;
	pinf := 1.0 / zero;
	fmt.Println(pinf);

	ninf := -1.0 / zero;
	fmt.Println(ninf);

	nan := zero / zero;
	fmt.Println(nan);

	/* 論理値型 */
	var t, fal bool = true, false;
	fmt.Println(t, fal);

	/* 文字列型 */
	var s_string string = "Hello Go";
	fmt.Println(s_string);
	fmt.Printf("%T\n", s_string);

	var si string = "300";
		fmt.Println(si);
	fmt.Printf("%T\n", si);
	
	//fmt.Println	そのまま表示する（改行付き）
// fmt.Printf	書式を指定して表示する（改行なし）

	fmt.Println(`test
			test
	test`);

	fmt.Println("\"");
	fmt.Println(`"`);

	fmt.Println(string(s_string[0]));

	/* byte(unit8)型 */
	byteA := []byte{72, 73};
	fmt.Println(byteA);
	fmt.Println(string(byteA));

	c := []byte("HI");
	fmt.Println(c);

	/* 配列型 */
	// 他の言語と違うのは後から要素数を変更できない

	var arr1 [3]int
	fmt.Println(arr1);
	fmt.Printf("%T\n", arr1);

	var arr2 [3]string = [3]string{"A", "B"};
	fmt.Println(arr2);

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3);

	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4);
	fmt.Printf("%T\n", arr4);

	fmt.Println(arr2[0]);
	fmt.Println(arr2[1]);
	fmt.Println(arr2[2]);
	fmt.Println(arr2[2-1]);

	arr2[2] = "C";
	fmt.Println(arr2);

	// arr5 [4]int
	// arr5 = arr1
	// fmt.Println(arr5);

	fmt.Println(len(arr1));

	/* interfaceとnil */
	// interface型は全ての型を内包できる特殊な型
	var x interface{}
	fmt.Println(x);

	x = 1
	fmt.Println(x);
		x = "Hello"
	fmt.Println(x);
		x = [3]int{1, 2, 3}
	fmt.Println(x);

	// x = 2
	// fmt.Printf(x+3);

	/* 型変換 */
	var i_change int = 1
	fl64 := float64(i_change)
	fmt.Println(fl64);
	fmt.Printf("i_change = %T\n", i_change);
	fmt.Printf("fl64 = %T\n", fl64);

	i2_change := int(fl64)
	fmt.Printf("i2_change = %T\n", i2_change);

	f1 := 10.5
i3 := int(f1)
fmt.Printf("i3 = %T\n", i3)
fmt.Println(i3)

var s string = "100"
fmt.Printf("s = %T\n", s)

// i, _ := strconv.Atoi(s)
// fmt.Println(i)
// fmt.Printf("i = %T\n", i)

var i2 int = 200
s2 := strconv.Itoa(i2)
fmt.Println(s2)
fmt.Printf("s2 = %T\n", s2)

var h string = "Hello World"
b := []byte(h)
fmt.Println(b)

h2 := string(b)
fmt.Println(h2)


}
package main

import "testing"

var Debug bool = false

func TestIsOne(t *testing.T) {
	// i := 0
	i := 1
	if Debug {
		t.Skip("スキップする")
	}
	v := IsOne(i)

	if !v {
		t.Errorf("%d != %d", i, 1)
	}
}

// pwdがgolang-webgosqlのときに go test -v ./...とすると、全部実行できる
// pwdがgolang-webgosqlのときに go test -cover ./...とすると、全部実行できる
// pwdがsection14/go_testのときに go test -v とすると、ここのテストだけ実行できる
// pwdがsection14/go_testのときに go test -cover とすると、ここのテストだけ実行できる

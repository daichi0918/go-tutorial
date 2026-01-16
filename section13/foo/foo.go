package foo

const (
	MAX = 100
	min = 1 // 最初の一文字目が小文字だとパッケージ外から参照できない
)

func ReturnMin() int {
	return min
}

package main
import "fmt"

// 型
// interface & nil
// interfaceはあらゆる型と互換性がある

func main() {
	// これだけだとnilが表示される
	var x interface{}
	fmt.Println(x)

	x = 1
	fmt.Println(x)

	// 値の更新はできるが計算はできない
	// fmt.Println(x + 3)とかはNG
	x = 3.14
	fmt.Println(x)
	x = "A"
	fmt.Println(x)
	x = [3]int{1, 2, 3}
	fmt.Println(x)
}
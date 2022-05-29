package main

import "fmt"

// スライス
// copy

func main() {
	sl := []int{100, 200}
	// これはコピーではなく、同じものを参照するので、更新するとslとsl2の両方が更新される
	sl2 := sl
	sl2[0] = 1000
	fmt.Println(sl) // 1000 200

	// 参照型でない変数は別々としてちゃんと扱われる
	var i int = 10
	i2 := i
	i2 = 100
	fmt.Println(i, i2) // 10 100


	sl := []int{1, 2, 3, 4, 5}
	sl2 := make([]int, 5, 10)
	n := copy(sl2, sl) // 返り値は要素数
	fmt.Println(n, sl2)
}

package main

import "fmt"

// スライス
// 宣言・操作

func main() {
	// スライスの宣言
	var sl []int
	fmt.Println(sl)

	// 値の格納、明示的宣言
	var sl2 []int = []int{100, 200}
	fmt.Println(sl2)

	// 暗黙的宣言
	sl3 := []string{"A", "B"}
	fmt.Println(sl3)

	// make関数で作成
	// 第一引数に型、第二引数に要素数
	sl4 := make([]int, 5)
	fmt.Println(sl4)

	// 値の代入、上書き
	sl2[0] = 1000
	fmt.Println(sl2)

	// 値の取り出し
	sl5 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl5)

	fmt.Println(sl5[0]) // 1

	fmt.Println(sl5[2:4]) // 3 4

	fmt.Println(sl5[:2]) // 1 2

	fmt.Println(sl5[2:]) // 3 4 5

	fmt.Println(sl5[:]) // 1 2 3 4 5

	fmt.Println(sl5[len(sl5)-1]) // 5

	fmt.Println(sl5[1 : len(sl5)-1]) // 2 3 4

}

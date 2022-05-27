package main

import "fmt"

// 定数
// 定数はグローバルスコープが多い
// 定数は上書きできない
// 頭文字を大文字にすると他パッケージでも参照できる
const Pi = 3.14

// 複数定義
const (
	URL = "https://aaa.com"
	SiteName = "test"
)

// 値の省略
// 一つ上の値が入るので1 1 1 A A Aと表示される
const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

// iota = 連続する数字 0 1 2と表示される
const (
	c0 = iota
	c1
	c2
)

func main() {
	fmt.Println(Pi)

	fmt.Println(URL, SiteName)

	fmt.Println(A, B, C, D, E, F)

	fmt.Println(c0, c1, c2)

}
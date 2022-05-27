package main

import "fmt"

/*
func 関数名(引数　引数の型) 戻り値型 {
	関数の中身
	return 返す値
}
*/

func Plus(x, y int) int {
	return x + y
}

// 戻り値が複数の場合
func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

// 戻り値の省略(戻り値の変数をあらかじめ定義した場合)
func Double(price int) (result int) {
	result = price * 2
	return
}

// 戻り値がない場合
func Noreturn() {
	fmt.Println("No Return")
	return
}

func main() {
	i := Plus(1, 2)
	fmt.Println(i)

	i2, _ := Div(9, 4)
	fmt.Println(i2)

	i3 := Double(100)
	fmt.Println(i3)

	Noreturn()

	// 無名関数
	f := func(x, y int) int {
		return x + y
	}
	f(1, 2)

	// 変数に代入しないパターンの無名関数
	// 簡易的な関数とかで利用する
	func(x, y int) int {
		return x + y
	}(1, 2)
}

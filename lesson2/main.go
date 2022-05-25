package main

import "fmt"

// 変数

// 関数外での暗黙定義はできない
// i5 := 600
// 明示的ならOK
var i5 int = 600

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

func main() {
	// 明示的な定義
	// var 変数名 型 = 値
	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t,f)

	var (
		i2 int = 200
		s2 string = "Golang"
	)
	fmt.Println(i2,s2)

	var i3 int
	var s3 string
	fmt.Println(i3,s3)
	i3 = 300
	s3 = "Go"

	i = 150
	fmt.Println(i)

	// 暗黙的な定義(型指定の必要がない)
	// 変数名 := 値
	i4 := 400
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)

	// 再定義はできない
	// i4 := 500
	// fmt.Println(i4)

	// 型推論でint型になっているため、string値は代入できない
	// i4 = "Hello"
	// fmt.Println(i4)

	fmt.Println(i5)

	outer()

	// Goでは変数を定義したら必ず使わなければならない
	// var s5 string = "Not Use"

	// 基本的には明示的定義が良い
	// 関数内なら暗黙定義でもOK
	// 関数外は明示的定義しかNG
}
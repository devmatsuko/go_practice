package main

import "fmt"

// 算術演算子
// 比較演算子

func main() {
	// 算術演算子
	fmt.Println(1 + 2)
	// 文字列の結合
	fmt.Println("ABC" + "DE")


	fmt.Println(5 - 1)

	fmt.Println(5 * 4)

	fmt.Println(60 / 3)

	fmt.Println(9 % 4)

	n := 0
	n += 4
	fmt.Println(n)
	n++
	fmt.Println(n)
	n--
	fmt.Println(n)

	s := "ABC"
	s += "DEF"
	fmt.Println(s)


	// 比較演算子
	// true
	fmt.Println(1 == 1)
	// false
	fmt.Println(1 == 2)
	// true
	fmt.Println(1 <= 8)
	// false
	fmt.Println(1 >= 8)
	

	// 論理演算子
	// false
	fmt.Println(true && false == true)
	// true
	fmt.Println(true && true == true)
	fmt.Println(true && false == false)

	// true
	fmt.Println(true || false == true)
}
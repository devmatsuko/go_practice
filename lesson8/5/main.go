package main

import "fmt"
// スライス 可変長引数(引数の数を固定しない)

// ...intはint型可変長引数
func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func main() {
	fmt.Println(Sum(1, 2, 3)) // 6

	fmt.Println(Sum(1, 2, 3, 4, 5, 6)) // 21

	fmt.Println(Sum())

	sl := []int{1, 2, 3}
	// スライスを渡す時はスライス変数...
	fmt.Println(Sum(sl...))
}

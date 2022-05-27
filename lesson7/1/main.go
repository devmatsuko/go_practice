package main

import "fmt"

// if
// 条件分岐

func main() {
	a := 1
	if a == 2 {
		fmt.Println("two")
	} else if a == 1 {
		fmt.Println("one")
	} else {
		fmt.Println("i don't know")
	}

	// 簡易文付
	// 変数のスコープはif文の中のみ
	if b := 100; b == 100 {
		fmt.Println("one hundred")
	}

	x := 0
	if x := 2; true {
		fmt.Println(x) // 2
	}
	fmt.Println(x) // 0
}

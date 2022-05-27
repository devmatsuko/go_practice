package main

import "fmt"

// スイッチ文

func main() {
	n := 1
	switch n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("I don't know")
	}

	// 簡易文を入れる場合
	switch n := 2; n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("I don't know")
	}

	// switch文の対象変数を省略する
	// case文に具体的な条件式がある場合のみOK
	n2 := 6
	switch {
	case n2 > 0 && n2 < 4:
		fmt.Println("0 < n2 < 4")
	case n2 > 3 && n2 < 7:
		fmt.Println("3 < n2 < 7")
	}

	// 変数と条件式が混在するのはNG
	/*
		switch n := 2; n {
		case 1, 2, 3:
			fmt.Println("0 < n2 < 4")
		case n > 3 && n < 6:
			fmt.Println("3 < n2 < 6")
		}
	*/

}

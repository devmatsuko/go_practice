package main

import "fmt"

// for文
// 繰り返し処理

func main() {

	i := 0
	for {
		i++
		if i == 3 {
			break
		}
		fmt.Println("Loop")
	}

	// 終了条件あり
	point := 0
	for point < 10 {
		fmt.Println(point)
		point++
	}

	// よく見るfor
	for i := 0; i < 100; i++ {
		if i == 3 {
			fmt.Println(3)
			continue
		}
		if i > 30 {
			fmt.Println("抜けた")
			break
		}
		fmt.Println(i)
	}

	// 配列とfor
	arr := [3]int{1, 2, 3}

	for i := 0; i < len(arr); i++ {
		fmt.Println(i, arr[i])
	}

	// スライスと配列
	sl := []string{"python", "PHP", "GO"}
	for i, v := range sl {
		fmt.Println(i, v)
	}

	// mapと配列
	m := map[string]int{"apple": 100, "banana": 200}

	for k, v := range m {
		fmt.Println(k, v)
	}
}

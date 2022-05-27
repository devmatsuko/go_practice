package main

import "fmt"

// ジェネレーター（クロージャーの応用）
// 何らかの法則に従って連続した値を返す

func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	ints := integers()

	fmt.Println(ints()) // => "1"
	fmt.Println(ints()) // => "2"
	fmt.Println(ints()) // => "3"

	// 再定義するとリセットされる
	otherInts := integers()
	fmt.Println(otherInts()) // => "1"
}

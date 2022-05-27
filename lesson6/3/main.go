package main

import "fmt"

// クロージャーの実装
// クロージャは一度定義すると中の変数情報が保持される
// クロージャーとは日本語では関数閉包と呼ばれ、関数と関数の処理に関する関数外の環境をセットして閉じ込めた物です。

func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func main() {
	f := Later()

	fmt.Println(f("Hello")) // => ""
	fmt.Println(f("My"))    // => "Hello"
	fmt.Println(f("name"))  // => "My"
}

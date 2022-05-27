package main

import "fmt"

// 関数を返す関数

// 関数を引数に取る関数
func callFunction(f func()) {
	f()
}

func ReturnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

func main() {
	f := ReturnFunc()
	f() // => "I'm a function"

	callFunction(func() {
		fmt.Println("I'm a function")
	}) // => "I'm a function"
}

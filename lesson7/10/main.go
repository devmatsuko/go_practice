package main

import "fmt"
// init
// 初期化

// mainの実行前に実行される
// initは複数定義できるが、あまり分けるメリットはない
func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("Main")
}

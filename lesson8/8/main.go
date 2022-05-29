package main

import "fmt"
// channel
// 複数のゴールーチン間でデータの受け渡しをするために設計されたデータ構造

func main() {
	// チャネルの宣言
	var ch1 chan int
	// 受信専用のチャネル
	// var ch2 <-chan int
	// 送信専用のチャネル
	// var ch3 chan<- int

	// make関数を使って初めて書き込み読み込みができるようになる
	ch1 = make(chan int)
	fmt.Println(cap(ch1)) // 0

	// 一気に宣言、makeを行う場合の書き方
	ch2 := make(chan int)
	fmt.Println(cap(ch2)) // 0

	// チャネルの容量(バッファサイズ)を定義する
	ch3 := make(chan int, 5)
	fmt.Println(cap(ch3)) // 5

	// ch3にデータを送信する
	ch3 <- 1
	fmt.Println(len(ch3)) // 1
	ch3 <- 2
	ch3 <- 3
	fmt.Println("len",len(ch3)) // 3

	// ch3からデータを受信する
	i := <-ch3
	fmt.Println(i) // 1(最初に入れたデータ)
	fmt.Println("len",len(ch3)) // 2
	i2 := <-ch3
	fmt.Println(i2) // 2(次に入れたデータ)
	fmt.Println(len(ch3)) // 1

	// 変数に入れずに受信データを表示する
	fmt.Println(<-ch3) // 3
}

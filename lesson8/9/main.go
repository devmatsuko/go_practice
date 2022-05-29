package main

import (
	"fmt"
	"time"
)

// チャネルとゴールーチン

func reciever(c chan int) {
	for {
		// 引数のチャネルから受信して表示する
		i := <-c
		fmt.Println(i)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	ch1 <- 1

	// レシーバを並行実行しておく
	// チャネルにデータが送信され次第、表示される
	go reciever(ch1)
	go reciever(ch2)

	i := 0
	for i < 100 {
		// チャネルにデータを送信する
		ch1 <- i
		ch2 <- i
		time.Sleep(50 * time.Millisecond)
		i++
	}
}

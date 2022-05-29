package main

import (
	"fmt"
	"time"
)
// チャネルのclose
// チャネルをmakeするとopen状態になり、明示的にcloseする

func reciever(name string, ch <-chan int) {
	for {
		i, ok := <-ch
		// チャネルがcloseかつチャネルが空の場合、breakする
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "END")
}

func main() {
	ch1 := make(chan int, 2)

	// ch1をクローズ
	close(ch1)

	//ch1 <- 1 // データの送信はクローズするのでエラー

	fmt.Println(<-ch1) // データの受信はできる

	i, ok := <-ch1
	fmt.Println(i, ok) // 0 false(チャネルが空かつcloseされている)

	go reciever("1.goroutin", ch1)
	go reciever("2.goroutin", ch1)
	go reciever("3.goroutin", ch1)

	i := 0
	for i < 100 {
		ch1 <- i
		i++
	}
	close(ch1)
	time.Sleep(3 * time.Second)
}

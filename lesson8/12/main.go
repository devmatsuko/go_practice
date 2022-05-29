package main

import (
	"fmt"
)
// チャネルのselect
// ゴールーチンを停止することなくデータを操作する

func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan string, 2)
	ch2 <- "A"
	//v1 := <-ch1
	//v2 := <-ch2
	//fmt.Println(v1)
	//fmt.Println(v2)

	select {
	case v1 := <-ch1: // ch1にデータが入った場合の処理
		fmt.Println(v1 + 1000)
	case v2 := <-ch2: // ch2にデータが入った場合の処理
		fmt.Println(v2 + "!?")
	default:
		fmt.Println("どちらでもない")
	}
}

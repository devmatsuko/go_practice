package main

import "fmt"
// チャネルのfor

func main() {
	ch3 := make(chan int, 3)
	ch3 <- 1
	ch3 <- 2
	ch3 <- 3
	close(ch3) // forの前にcloseを使わないと、空のチャネルを参照してしまい、deadlockとなる
	for i := range ch3 {
		fmt.Println(i)
	}
}

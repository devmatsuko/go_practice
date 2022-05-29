package main

import (
	"fmt"
	"time"
)

// goルーチンによる並行処理

func sub() {

	for {
		fmt.Println("Sub Loop")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// goをつけることでゴールーチンとして並行実行できる
	go sub()

	for {
		fmt.Println(("Main Loop"))
		time.Sleep(200 * time.Millisecond)
	}
}

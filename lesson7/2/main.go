package main

import (
	"fmt"
	"strconv"
)

// エラーハンドリング

func main() {
	var s string = "100"
	// errを返す関数を呼ぶ
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Printf("i = %T\n", i)
	}
}

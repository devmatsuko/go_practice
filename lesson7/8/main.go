package main

import "fmt"

// panic&recover
// プログラムを強制終了させる
// なるべくエラー処理で済ませる方が望ましい

func main() {
	defer func() {
		// panicで中断したプログラムを復活させる
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	// panicを用いてプログラムを強制終了させる
	panic("rungime error")
	fmt.Println("Start")

}

package main

import "fmt"

// ラベル付きfor

func main() {
	Loop:
		for {
			for {
				for {
					fmt.Println("Start")
					break Loop
				}
				fmt.Println("処理をしない")
			}
			fmt.Println("処理をしない")
		}
		fmt.Println("End")
}
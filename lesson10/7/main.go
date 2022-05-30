package main

import "fmt"
// 独自型

// 独自のintを作る
type MyInt int

// レシーバをMyInt型に固定できる
func (mi MyInt) Print() {
	fmt.Println(mi)
}

func main() {
	var mi MyInt
	fmt.Println(mi) // 0
	fmt.Printf("%T\n", mi) // main.MyInt

	//i := 100
	//fmt.Println(mi + i) // MyInt型とint型では計算できない。他のint型と計算できないように縛る際に使える。

	mi.Print()
}

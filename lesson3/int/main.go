package main
import "fmt"

// 型
// 整数型

func main() {
	// int はint8~int64が存在し、何も指定しない場合はPC依存で決まる
	var i int = 100

	var i2 int64 = 200
	fmt.Println(i + 50)

	// 同じintでも型が違うと計算できない
	// fmt.Println(i + i2)

	// 現在の型を調べる。%Tは変数の型を表示する
	fmt.Printf("%T\n", i)

	// 型をint32に変換する
	fmt.Printf("%T\n", int32(i2))
}
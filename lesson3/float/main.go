package main
import "fmt"

// 型
// 浮動小数点型

func main() {
	var fl64 float64 = 2.4
	fmt.Println(fl64)

	// 暗黙定義ではfloat64になる
	fl := 3.2
	fmt.Println(fl + fl64)
	fmt.Printf("%T, %T\n", fl64, fl)

	// float32はあまり使わない
	var fl32 float32 = 1.2
	fmt.Printf("%T\n", fl32)

	// 型変換
	fmt.Printf("%T\n", float64(fl32))
}
package main
import "fmt"

// 型
// 文字列型

func main() {
	var s string = "Hello, String"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	var si string = "300"
	fmt.Println(si)
	fmt.Printf("%T\n", si)

	// 複数行にまたがる表示
	fmt.Println(`test
	test
	test`)

	// ダブルクォーテーションを文字列として表示する \を使うor``を使う
	fmt.Println("\"")
	fmt.Println(`"`)
}
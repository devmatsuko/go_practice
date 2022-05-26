package main
import "fmt"

// 型
// 配列型
// 基本的に他の言語と違って要素数が変えられない
// 要素数を変える場合はスライスを使う
// 配列型　＝　要素数を変更できない。
// スライス型　＝　要素数を変更可能。

func main() {
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	var arr2 [3]string =[3]string{"A", "B"}
	fmt.Println(arr2)

	// 暗黙定義
	arr3 := [3]int{1,2,3}
	fmt.Println(arr3)

	// 要素数の省略(自動でカウントされる)
	arr4 := [...]string{"C","D"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	// 値の取り出し
	fmt.Println(arr2[0])

	// 値の更新
	arr2[2] = "C"
	fmt.Println(arr2)

	// 要素数が異なる配列を代入することはできない
	// cannot use arr1 (variable of type [3]int) as type [4]int in assignment
	// var arr5 [4]int
	// arr5 = arr1
	// fmt.Println(arr5)

	// 要素数を調べる関数
	fmt.Println(len(arr1))
}
package main

import (
	"fmt"
)
// ポインタ
// メモリ上のアドレスの型と値の情報

func Double(i int) {
	i = i * 2
}

func Doublev2(i *int) {
	*i = *i * 2
}

func main() {
	var n int = 100
	fmt.Println(n) // 100

	fmt.Println(&n) // アドレス番地の表示。0xc000016070

	Double(n) // n自体が倍にするのではなく、nをコピーしたiを2倍にしているだけ
	fmt.Println(n) // 100

	// ポインタの宣言
	var p *int = &n
	fmt.Println(p) // アドレス番地の表示。0xc000016070
	fmt.Println(*p) // 100。アドレス番地の値を表示。実体。

	Doublev2(&n)
	fmt.Println(n) // 200。ポインタ型で引数を渡すことで、コピーではなく実体を倍にしてくれる。

	Doublev2(p)
	fmt.Println(*p) // 400。上と同じ
}

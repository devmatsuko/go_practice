package main

import "fmt"
// マップ

func main() {
	// 明示的宣言 map[キー型]バリュー型
	var m = map[string]int{"A": 100, "B": 200}
	fmt.Println(m)

	// 暗黙的宣言
	m2 := map[string]int{"A": 100, "B": 200}
	fmt.Println(m2)

	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3)

	// make関数での作成
	m4 := make(map[int]string)
	m4[1] = "JAPAN"
	m4[2] = "USA"
	m4[3] = "CHINA"
	fmt.Println(m4)

	// 値の取り出し
	fmt.Println(m["A"])
	fmt.Println(m4[2])
	fmt.Println(m4[3]) // 登録されていない場合、空文字が帰ってくる

	// 二つ目の引数には取り出しに成功したかのフラグが帰ってくる
	s, ok := m4[1]
	fmt.Println(s, ok)
	if !ok {
		fmt.Println("error")
	}

	// mapから要素を削除
	// delete(変数名, キー)
	delete(m4, 3)
	fmt.Println(m4)
}

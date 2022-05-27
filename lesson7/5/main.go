package main

import "fmt"

// 型スイッチ

// 全ての型と互換性があるinterface型を引数に取る関数
func anything(a interface{}) {
	//fmt.Println(a)
	switch v := a.(type) {
	case string:
		fmt.Println(v + "!?")
	case int:
		fmt.Println(v + 40000)
	}
}

func main() {
	anything("aaa")
	anything(1)

	var x interface{} = 3
	// interface型をintとして復元する
	// これをしないとint型と計算ができない
	i := x.(int)
	fmt.Println(i)

	// xはintとして認識されているため、floatとして復元できない
	f, isFloat64 := x.(float64)
	// 0 falseが表示される
	fmt.Println(f, isFloat64)

	if x == nil {
		fmt.Println("None")
	} else if i, isInt := x.(int); isInt { // int型参照が成功した場合
		fmt.Println(i, "x is int")
	} else if s, isString := x.(string); isString { // string型参照が成功した場合
		fmt.Println(s, isString)
	} else {
		fmt.Println("I don't know")
	}

	// スイッチ文で型参照ごとに処理を分ける
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("I don't Know")
	}

	switch v := x.(type) {
	case bool:
		fmt.Println(v, "bool")
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	default:
		fmt.Println("I don't Know")
	}
}

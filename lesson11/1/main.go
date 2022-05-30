package main

import "fmt"
// インタフェース
// 最もポピュラーな使い方。異なる型に共通の性質を付与する

type Stringify interface {
	// ToStringメソッドは文字列型を返す性質を与える
	ToString() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("Name=%v, Age=%v", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("Number=%v, Model=%v", c.Number, c.Model)
}

func main() {
	// personとcarをStringify型という同じ型としてまとめる
	vs := []Stringify{
		&Person{Name: "Taro", Age: 15},
		&Car{Number: "XXX-012", Model: "RX-78"},
	}

	for _, v := range vs {
		fmt.Println(v.ToString())
	}
}

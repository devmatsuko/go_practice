package main

import "fmt"
// 構造体のメソッド

type User struct {
	Name string
	Age  int
	//X, Y int
}

// ()の中はレシーバ。構造体(モデル?)変数が入る
// 構造体のメソッドになるため、u.SayName()で呼び出せる
func (u User) SayName() {
	fmt.Println(u.Name)
}

// レシーバは参照型なので値は更新されない。厳密にはコピーが更新されるだけ
func (u User) SetName(name string) {
	u.Name = name
}

// レシーバをポインタ型にすることで実体を更新することができる
func (u *User) SetName2(name string) {
	u.Name = name
}

func main() {
	user1 := User{Name: "user1"}
	user1.SayName() // user1

	user1.SetName("A")
	user1.SayName() // user1。参照型なので更新はされていないため。

	user1.SetName2("A") // ポインタ型で更新する
	user1.SayName() // A

	user2 := &User{Name: "user2"}
	user2.SetName2("B")
	user2.SayName()
}

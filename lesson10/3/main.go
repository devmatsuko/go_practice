package main

import "fmt"

// 構造体埋め込み
// フィールドに構造体を埋め込むことができる

type T struct {
	User User // UserフィールドにUser構造体を埋め込める
}

type User struct {
	Name string
	Age  int
	//X, Y int
}

func (u *User) SetName() {
	u.Name = "A"
}

func main() {
	t := T{User: User{Name: "user1", Age: 10}}

	fmt.Println(t) // {{user1 10}}

	fmt.Println(t.User) // {user1 10}

	fmt.Println(t.User.Name) // user1

	//fmt.Println(t.Name)

	t.User.SetName() // 構造体メソッドの呼びだし。構造体埋め込みの場合は頭にt.が必要
	fmt.Println(t) // // {{A 10}}
}

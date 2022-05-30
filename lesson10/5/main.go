package main

import "fmt"
// 構造体とスライス

type User struct {
	Name string
	Age  int
	//X, Y int
}

// Userポインタ型のスライス
type Users []*User

// ↑の方が望ましい
/*
type Users struct {
	Users []*Users
}
*/

func main() {
	user1 := User{Name: "user1", Age: 10}
	user2 := User{Name: "user2", Age: 20}
	user3 := User{Name: "user3", Age: 30}
	user4 := User{Name: "user4", Age: 40}

	// Users型の変数の定義。
	users := Users{}

	// スライスにuserを追加
	users = append(users, &user1)
	users = append(users, &user2)
	users = append(users, &user3, &user4)

	fmt.Println(users) // [0x... , 0x..., 0x..., 0x...]

	for _, u := range users {
		fmt.Println(*u) // 各userが表示される
	}

	// make関数でも作れる
	users2 := make([]*User, 0)
	users2 = append(users2, &user1, &user2)

	for _, u := range users2 {
		fmt.Println(*u)
	}
}

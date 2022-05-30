package main

import "fmt"

// 構造体の宣言
type User struct {
	Name string
	Age  int
	// X, Y int
}

func UpdateUser(user User) {
	user.Name = "A"
	user.Age = 1000
}

// 引数がポインタ型
func UpdateUser2(user *User) {
	user.Name = "A"
	user.Age = 1000
}

func main() {
	// 構造体User型の変数の定義
	var user1 User
	fmt.Println(user1) // { 0}
	user1.Name = "user1"
	user1.Age = 10
	fmt.Println(user1) // {user1 10}

	// 暗黙宣言
	user2 := User{}
	fmt.Println(user2) // { 0}
	user2.Name = "user2"
	fmt.Println(user2) // {user2 0}

	user3 := User{Name: "user3", Age: 30}
	fmt.Println(user3) // {user3 30}

	// フィールド名を指定しない場合は、順序が合っていればOK
	user4 := User{"user4", 40}
	fmt.Println(user4) // {user4 40}

	// 下記はNG
	//user5 := User{30, "user5"}
	//fmt.Println(user5)

	user6 := User{Name: "user6"}
	fmt.Println(user6) // {user6 0}

	// newを使った変数定義。ポインタ型で定義される
	user7 := new(User)
	fmt.Println(user7) // &{ 0}

	// アドレス演算子を使って定義。newと同じ動きになる
	user8 := &User{}
	fmt.Println(user8) // &{ 0}

	// 引数が普通
	UpdateUser(user1)
	// 引数がポインタ型
	UpdateUser2(user8)

	fmt.Println(user1) // {user1 10} 参照型なので更新されない
	fmt.Println(user8) // &{A 1000} 値型なので更新される
}

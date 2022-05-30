package main

import "fmt"

// 構造体コンストラクタ
// Goにはコンストラクはないが、擬似的に作れる

type User struct {
	Name string
	Age  int
	//X, Y int
}

// コンストラクタ
func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

func main() {
	// コンストラクタ実行
	user1 := NewUser("user1", 10)

	fmt.Println(user1) // &{user1 10}

	fmt.Println(*user1) // {user1 10}
}

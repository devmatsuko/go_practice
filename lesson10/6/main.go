package main

import "fmt"

// マップと構造体

type User struct {
	Name string
	Age  int
	//X, Y int
}

func main() {
	// キーがint型、バリューがUser型のマップの定義
	m := map[int]User{
		1: {Name: "user1", Age: 10},
		2: {Name: "user2", Age: 20},
	}

	fmt.Println(m)

	// キーがUser型、バリューがstring型のマップの定義
	m2 := map[User]string{
		{Name: "user1", Age: 10}: "Tokyo",
		{Name: "user2", Age: 20}: "LA",
	}

	fmt.Println(m2)

	// make関数での定義
	m3 := make(map[int]User)
	fmt.Println(m3)
	m3[1] = User{Name: "user3"}
	fmt.Println(m3)

	for _, v := range m {
		fmt.Println(v)
	}
}

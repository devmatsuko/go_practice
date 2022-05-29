package main

import "fmt"
// スライスのfor

func main() {
	sl := []string{"A", "B", "C"}
	fmt.Println(sl)

	// iはindex、vはvalue
	for i, v := range sl {
		fmt.Println(i, v)
	}

	/*
	  for _, v := range sl {
			fmt.Println(v)
	  }
	*/
	/*
		for i := range sl {
			fmt.Println(i)
		}
	*/

	for i := 0; i < len(sl); i++ {
		fmt.Println(i, sl[i])
	}
}

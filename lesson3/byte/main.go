package main
import "fmt"

// 型
// byte型(uint8)

func main() {
	byteA := []byte{72, 73}
	fmt.Println(byteA)

	// 文字列に変換。アスキーコードが呼び出されるため、HIが表示される
	fmt.Println(string(byteA))

	// 文字列をbyte型に変換 HI → [72 73]
	c := []byte("HI")
	fmt.Println(c)
}
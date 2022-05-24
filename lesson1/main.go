// mainパッケージに関するプログラム
// パッケージの宣言は1つのみ
package main

// フォーマットパッケージのインポート
import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(time.Now())
}
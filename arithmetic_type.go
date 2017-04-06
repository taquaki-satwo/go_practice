package main

import "fmt"

// C1 コメント
const C1 = 12345

// C2 コメント
const C2 int = 34567

func main() {
	// var x int = C1 + C2
	var x = C1 + C2

	fmt.Println("12345 + 3456 =", x)

	var a int32 = 123
	var b int64 = 234

	fmt.Println("123 + 234 =", a+int32(b))
}

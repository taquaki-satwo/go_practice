package main

import "fmt"

// MyError コメント
type MyError string

func main() {
	var i interface{} = 123

	// var s int = i.(int)
	var s = i.(int)

	fmt.Println(s)
}

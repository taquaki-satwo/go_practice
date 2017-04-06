package main

import "fmt"

type myType int

func main() {
	fmt.Println("開始")

	defer delay()

	fmt.Println("終了")
}

func delay() {
	fmt.Println("deley")
}

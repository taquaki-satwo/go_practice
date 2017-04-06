package main

import "fmt"

type MyError string

func (e MyError) Error() string {
	return string(e)
}

func main() {
	err := do()

	if err == nil {
		fmt.Println("エラーなし")
	} else {
		fmt.Println("エラーあり??", err)
	}
}

func do() error {
	var ret *MyError = nil

	return ret
}

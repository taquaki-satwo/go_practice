package main

import (
	"fmt"
	"strconv"
)

func fizzbuzz(num int) string {
	for i := 0; i <= num; i++ {
		fizz := i % 3
		buzz := i % 5
		switch {
		case fizz == 0 && buzz == 0:
			fmt.Println("fizzbuzz")
		case fizz == 0 && buzz != 0:
			fmt.Println("fizz")
		case fizz != 0 && buzz == 0:
			fmt.Println("buzz")
		default:
			fmt.Println(strconv.Itoa(i))
		}
	}
	return "\nEnd"
}

func main() {
	fmt.Println(fizzbuzz(100))
}

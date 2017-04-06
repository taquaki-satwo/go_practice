package main

import "fmt"

func fib(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(fib(i))
	}
}

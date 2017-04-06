package main

import "fmt"

func main() {
	var ptr *int

	var i int = 12345

	ptr = &i

	fmt.Println("i adress:", &i)
	fmt.Println("ptr value:", ptr)

	fmt.Println("i value:", i)
	fmt.Println("Pointer i value:", *ptr)

	*ptr = 999

	fmt.Println("Pointer change i value:", i)
}

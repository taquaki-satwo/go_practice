package main

import "fmt"

func main() {
	dst := []int{1, 2, 3, 4}
	src := []int{5, 6, 7}

	count := copy(dst[1:], src)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println("count:", count)
}

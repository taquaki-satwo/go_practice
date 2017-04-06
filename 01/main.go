package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start")

	for i := 1; i < 100; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}

	fmt.Println("Stop")
}

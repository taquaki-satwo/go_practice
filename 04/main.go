package main

import (
	"fmt"
)

func main() {
	result := cal(1, 2, 3, 4)

	if result < 2100 {
		fmt.Println("Laminar flow")
	} else if result < 4000 {
		fmt.Println("Transient flow")
	} else {
		fmt.Println("Turbulent flow")
	}
}

func cal(d int, v int, rho int, mu int) int {
	return d * v * rho / mu
}

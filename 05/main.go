package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// result := cal(1, 2, 3, 4)

	for {
		fmt.Println("Do you want to calculate again (y/n)")
		var res string
		fmt.Scan(&res)

		if res == "y" {
			scanner := bufio.NewScanner(os.Stdin)
			stdin := [4]int{}

			for scanner.Scan() {
				stdin = append(stdin, scanner.Text())
				fmt.Println(stdin)
			}
		} else if res == "n" {
			os.Exit(0)
		} else {
			continue
		}
	}

	// if result < 2100 {
	// 	fmt.Println("Laminar flow")
	// } else if result < 4000 {
	// 	fmt.Println("Transient flow")
	// } else {
	// 	fmt.Println("Turbulent flow")
	// }
}

// func cal(d int, v int, rho int, mu int) int {
// 	return d * v * rho / mu
// }

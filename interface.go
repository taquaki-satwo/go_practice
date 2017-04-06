package main

import "fmt"

type Calculator interface {
	Calculator(a int, b int) int
}

type Add struct {
}

func (x Add) Calculator(a int, b int) int {
	return a + b
}

type Sub struct {
}

func (x Sub) Calculator(a int, b int) int {
	return a - b
}

func main() {
	var add Add
	var sub Sub

	var cal Calculator

	cal = add

	fmt.Println("和:", cal.Calculator(1, 2))

	cal = sub

	fmt.Println("差:", cal.Calculator(1, 2))
}

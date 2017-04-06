package main

import "fmt"

type embedded struct {
	i int
}

func (x embedded) doSomething() {
	fmt.Println("test.doSomething()")
}

type test struct {
	embedded
}

func main() {
	var x test

	x.doSomething()
	x.i = 2
	fmt.Println(x.i)
}

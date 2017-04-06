package main

import "fmt"

func main() {
	var fname, lname string
	fmt.Println("firstname lastname")
	fmt.Scanln(&fname, &lname)

	var age int
	fmt.Println("age")
	fmt.Scanf("%d", &age)

	fmt.Printf("name: %s %s age: %d\n", fname, lname, age)
}

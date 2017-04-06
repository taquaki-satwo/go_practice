package main

import "fmt"

func main() {
  x := [5]string{"a", "b", "c", "d", "e"}

  var s1 []string

  s1 = x[:]
  fmt.Println(s1)
}
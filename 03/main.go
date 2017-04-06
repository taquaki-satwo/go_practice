package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	stdin := []string{}

	for scanner.Scan() {
		stdin = append(stdin, scanner.Text())

		sort.Strings(stdin)
		fmt.Println(stdin)
	}
}

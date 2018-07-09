package main

import (
	"fmt"
)

func sum(x int, y int) int {
	return x + y
}

func main() {
	var c = sum(5, 5)
	fmt.Println(c)
}

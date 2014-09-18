package main

import (
	"fmt"
)

var (
	a int = 1
	b int = 2
)

func main() {
	fmt.Println(a, b)

	var array [2]string
	array[0] = "new"
	array[1] = "egg"
	fmt.Println(array)
}

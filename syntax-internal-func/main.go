package main

import (
	"fmt"
)

func main() {
	total := func(x, y int) int {
		return x + y
	}

	fmt.Println("100+10=", total(100, 10))
}

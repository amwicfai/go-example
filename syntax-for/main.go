package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i > 6 {
			fmt.Println(i * 10)
		}
	}
}

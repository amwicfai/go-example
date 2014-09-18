package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("1")
	print()
	fmt.Println("5")
}

func print() {
	defer fmt.Println("2")
	defer fmt.Println("3")

	fmt.Println("4")
}

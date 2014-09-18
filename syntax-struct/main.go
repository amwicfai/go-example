package main

import "fmt"

type Size struct {
	Width  int
	Height int
}

func main() {
	fmt.Println(Size{Height: 1, Width: 2})

	size := Size{}
	size.Height = 100
	size.Width = 200

	fmt.Println(size.Height)

	size_pointer := &size
	size_pointer.Height = 1000
	fmt.Println(size.Height)
}

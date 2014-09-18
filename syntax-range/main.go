package main

import (
	"fmt"
)

func main() {
	var a []int
	a = append(a, 1, 2, 3)
	fmt.Println(a)

	for i, v := range a {
		fmt.Printf("a[%d]:%d\r\n", i, v)
	}

	for i := range a {
		fmt.Printf("a[%d]\r\n", i)
	}

}

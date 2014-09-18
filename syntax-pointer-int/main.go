package main

import (
	"fmt"
)

func main() {
	i, j := 42, 9
	p := &i

	fmt.Println(*p)

	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 3
	fmt.Println(j)
}

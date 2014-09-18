package main

import (
	"fmt"
	"strconv"
)

var c, python, java bool

var z string = "100"

func main() {
	var i int
	x, y := "xxx", "yyy"

	zz, err := strconv.Atoi(z)
	if err != nil {
		fmt.Println(err)
		return
	}
	//z = string(zz)

	fmt.Println(x, y, zz, i, c, python, java)
}

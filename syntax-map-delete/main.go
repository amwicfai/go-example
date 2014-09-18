package main

import (
	"fmt"
)

var m map[string]int

func main() {
	m = make(map[string]int)
	m["zeeman"] = 12
	m["james"] = 22
	m["vincent"] = 32

	show_map()

	delete(m, "vincent")

	show_map()

	v, ok := m["vincent"]
	fmt.Println("The value:", v, "IsExists?", ok)

	v, ok = m["james"]
	fmt.Println("The value:", v, "IsExists?", ok)
}

func show_map() {
	fmt.Println("===============")
	for name, age := range m {
		fmt.Printf("%s:%d\r\n", name, age)
	}
}

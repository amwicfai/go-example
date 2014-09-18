package main

import (
	"fmt"
)

type Order struct {
	Id int
}

func main() {
	orders := []Order{Order{1}, Order{2}, Order{3}}
	fmt.Println(orders)

	// 始终为集合中对应索引的值拷贝，因此它一般只具有只读性质，对它所做的任何修改都不会影响到集合中原有的值
	for _, order := range orders {
		order.Id = order.Id * 10
	}
	fmt.Println(orders)

	for i := range orders {
		orders[i].Id = orders[i].Id * 10
	}
	fmt.Println(orders)
}

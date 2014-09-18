package main

import (
	"fmt"
)

func main() {
	id, name := GetUser()
	fmt.Println(id, name)

	/*
	   空白标识符 _ 也被用于抛弃值，如值 5 在：_, b = 5, 7 中被抛弃。
	   _ 实际上是一个只写变量，你不能得到它的值。
	   	这样做是因为 Go 语言中你必须使用所有被声明的变量，
	   	但有时你并不需要使用从一个函数得到的所有返回值。
	*/
	_, name2 := GetUser()
	fmt.Println(name2)
}

func GetUser() (id int, name string) {
	return 100, "zeeman"
}

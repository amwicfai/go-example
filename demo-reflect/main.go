package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	name string
}

func (this *MyStruct) GetName() string {
	return this.name
}

func main() {
	fmt.Println("---------------")

	var a MyStruct
	b := new(MyStruct)
	fmt.Println(reflect.ValueOf(a))
	fmt.Println(reflect.ValueOf(b))

	fmt.Println("---------------")
	a.name = "zeeman"
	b.name = "zeeman"
	val := reflect.ValueOf(a).FieldByName("name")

	fmt.Println(val)

	fmt.Println("---------------")
	fmt.Println(reflect.ValueOf(a).FieldByName("name").CanSet())
	fmt.Println(reflect.ValueOf(&(a.name)).Elem().CanSet())

	fmt.Println("---------------")
	var c string = "zeeman"
	p := reflect.ValueOf(&c)
	fmt.Println(p.CanSet())
	fmt.Println(p.Elem().CanSet())
	p.Elem().SetString("zeeman2")
	fmt.Println(c)
}

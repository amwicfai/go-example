//http://golang.org/ref/spec#Conversions

package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 1
	fmt.Print(reflect.Typeof(a))
}

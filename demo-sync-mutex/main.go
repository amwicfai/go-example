package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex
var a string

func main() {
	mutex.Lock()
	go foo()

	//阻塞直到unlock被调用
	mutex.Lock()
	fmt.Println(a)
}

func foo() {
	a = "hello,world"
	time.Sleep(3 * time.Second)
	mutex.Unlock()
}

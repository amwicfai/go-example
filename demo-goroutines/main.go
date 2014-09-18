package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world1")
	go say("world2")
	say("hello1")
	say("hello2")
	fmt.Println("begin")
}

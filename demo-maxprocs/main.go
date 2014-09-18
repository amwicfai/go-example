package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	go sheep(1)
	go sheep(2)
	time.Sleep(time.Millisecond)
}
func sheep(i int) {
	for ; ; i += 2 {
		fmt.Println(i, "只羊")
	}
}

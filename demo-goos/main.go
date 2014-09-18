package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Print("Go runs on ")
	os := runtime.GOOS
	switch os {
	case "linux":
		fmt.Println("Linux")
	default:
		//注释
		fmt.Printf("%s.", os)
	}

	today := time.Now()
	fmt.Println(today)
}

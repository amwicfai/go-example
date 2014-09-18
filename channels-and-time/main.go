package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	//当 select 中的其他条件分支都没有准备好的时候，`default` 分支会被执行。
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("...")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

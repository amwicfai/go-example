package main

import (
	"fmt"
	"time"
)

func main() {
	//计算时间
	i1 := time.Now()
	time.Sleep(500 * time.Millisecond)
	i2 := time.Now()
	i3 := i2.Sub(i1)
	took := int(i3.Seconds() * 1000)
	fmt.Printf("it took %dms\n", took)

	//使用变量sleep
	sleepms := 1000
	time.Sleep(time.Duration(sleepms) * time.Millisecond)

	//时间格式化, 注意不是yyyy-mm-dd这样的格式
	now := time.Now().Format("2006-01-02T15:04:05Z")
	fmt.Print(now)
}

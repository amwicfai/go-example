package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go fireError(3)
	go doWork()

	time.Sleep(time.Second * 10)
}

//每个协程都要自己捕获错误，否则异常导致进程退出
func fireError(index int) {
	defer func() {
		// 这里的err其实就是panic传入的内容
		if err := recover(); err != nil {
			//超过1024被截断，对于异常判断，已经足矣
			trace := make([]byte, 1024)
			//如果选择输出所有 goroutine 的 traceback, 会挂起所有goroutine
			count := runtime.Stack(trace, false)
			fmt.Printf("Uncaught exception: %s\nStack of %d bytes: %s", err, count, trace)
		}
	}()

	arr := []int{1, 2, 3}
	fmt.Println(arr[index])
}

func doWork() {
	fmt.Println("working...")
}

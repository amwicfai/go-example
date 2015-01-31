package main

import (
	"errors"
	"fmt"
	"runtime"
	"time"
)

func handlerException() {
	// 这里的err其实就是panic传入的内容
	fmt.Println("...")
	if err := recover(); err != nil {
		//超过1024被截断，对于异常判断，已经足矣
		trace := make([]byte, 1024)
		//如果选择输出所有 goroutine 的 traceback, 会挂起所有goroutine
		count := runtime.Stack(trace, false)
		fmt.Printf("Uncaught exception: %s\nStack of %d bytes: %s", err, count, trace)
	}
}

func main() {
	defer handlerException()

	//发生致命错误
	//panic("xxx")

	go fireError(3)
	//go doWork()

	time.Sleep(time.Second * 3)
}

//每个协程都要自己捕获错误，否则异常导致进程退出
func fireError(index int) {
	//panic异常为致使错误，协程内部panic异常不处理会导致进程崩溃
	defer handlerException()

	panic("yyy")

	arr := []int{1, 2, 3}
	fmt.Println(arr[index])
}

func doWork() error {
	err := errors.New("emit macho dwarf: elf header corrupted")
	return err
	//if err != nil {
	//	fmt.Print(err)
	//}
	//fmt.Println("working...")
}

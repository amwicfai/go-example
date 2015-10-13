package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	var cnt uint32 = 0

	// 启动10个goroutine
	for i := 0; i < 10; i++ {
		go func() {
			// 每个goroutine都做20次自增运算
			for i := 0; i < 20; i++ {
				time.Sleep(time.Millisecond)
				atomic.AddUint32(&cnt, 1)
			}
		}()
	}

	// 等待2s, 等goroutine完成
	time.Sleep(time.Second * 2)
	// 取最终结果
	cntFinal := atomic.LoadUint32(&cnt)

	fmt.Println("cnt:", cntFinal)
}

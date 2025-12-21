// Package concurrency 03_并发编程
package concurrency

import (
	"fmt"
	"time"
)

// select 能在 Channel 上进行非阻塞的收发操作
// select 在遇到多个 Channel 同时响应时，会随机执行一种情况

// select 非阻塞收发
func unbuffered() {
	ch := make(chan int)
	select {
	case i := <-ch:
		fmt.Println(i)
	default:
		fmt.Println("default")
	}
}

// 随机执行
func randDo() {
	ch := make(chan int)
	go func() {
		for range time.Tick(1 * time.Second) {
			ch <- 0
		}
	}()

	// 方便测度，这里加一个10秒的timer
	timer := time.NewTimer(10 * time.Second)

	for {
		select {
		case <-ch:
			fmt.Println("case1")
		case <-ch:
			fmt.Println("case2")
		case <-timer.C:
			fmt.Println("stop randDo")
			return
		}
	}
}

func Concurrency() {
	unbuffered()
	randDo()
	syncMutex()
	syncRWMutex()
	syncWaitGroup()
	syncOnce()
	syncCond()
}

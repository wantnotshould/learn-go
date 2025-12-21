// Package concurrency 03_并发编程练习题
package concurrency

import (
	"fmt"
	"sync"
)

func toCount(val *int, m *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	m.Lock()
	for range 1000 {
		*val++
	}
	m.Unlock()
}

// 计数器
func exercises1() {
	result := 0
	var m sync.Mutex
	var wg sync.WaitGroup

	// 启动 10 个 goroutine 来增加计数器
	for range 10 {
		wg.Add(1)
		go toCount(&result, &m, &wg)
	}

	wg.Wait()

	fmt.Println("result:", result)
}

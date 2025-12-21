// Package concurrency 03_并发编程
package concurrency

import (
	"fmt"
	"sync"
)

// sync.Mutex 互斥锁
// 用于对共享资源加锁，避免数据竞争

var mu sync.Mutex
var counter int

func increment() {
	mu.Lock() // 锁定资源
	counter++
	mu.Unlock() // 解锁资源
}

// 注：wg.Add wg.Done 须成对使用
func syncMutex() {
	var wg sync.WaitGroup

	// 启动多个goroutine
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()

		// // go1.25+
		// wg.Go(func() {
		// 	increment()
		// })
	}

	wg.Wait()
	fmt.Println("final counter value:", counter)
}

// sync.RWMutex 读写锁
var rwMutex sync.RWMutex
var counter2 int

func read() int {
	rwMutex.RLock()         // 加读锁
	defer rwMutex.RUnlock() // 解读锁
	return counter2
}

func write(val int) {
	rwMutex.Lock()         // 加写锁
	defer rwMutex.Unlock() // 解写锁
	counter2 = val
}

func syncRWMutex() {
	var wg sync.WaitGroup

	// 启动多个读写操作
	// 新写法
	for i := range 10 {
		wg.Go(func() {
			write(i) // 写操作
		})
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("read", read())
		}()
	}

	wg.Wait()
}

// sync.WaitGroup 等待组

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("task %d is completed\n", id)
}

func syncWaitGroup() {
	var wg sync.WaitGroup
	// 启动 5 个 goroutine
	for i := 1; i <= 5; i++ {
		wg.Add(1) // 增加等待的任务数
		go task(i, &wg)
	}

	wg.Wait() // 等待所有任务完成
	fmt.Println("all tasks are done.")
}

// sync.Once 单次操作
var once sync.Once

func initOnce() {
	fmt.Println("Initializing...")
}

func syncOnce() {
	// 同一操作仅会执行一次
	for i := 0; i < 10; i++ {
		go once.Do(initOnce)
	}

	// 等待 goroutine 执行完
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		once.Do(initOnce)
	}()
	wg.Wait()
}

// sync.Cond 条件变量

var (
	m        = sync.Mutex{}
	cond     = sync.NewCond(&m)
	counter3 = 0
)

func producer() {
	for {
		cond.L.Lock() // 加锁
		if counter3 >= 10 {
			cond.Wait() // 等待条件满足
		}

		counter3++
		fmt.Println("produced:", counter3)
		cond.Signal()   // 通知消费者
		cond.L.Unlock() // 解锁
	}
}

func consumer() {
	for {
		cond.L.Lock() // 加锁
		if counter3 == 0 {
			cond.Wait() // 等待条件满足
		}
		fmt.Println("consumed:", counter3)
		counter3--
		cond.Signal()   // 通知生产者
		cond.L.Unlock() // 解锁
	}
}

// TODO:10秒后停止下
func syncCond() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		producer()
	}()

	go func() {
		defer wg.Done()
		consumer()
	}()

	wg.Wait()
}

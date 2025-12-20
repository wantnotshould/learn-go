// Package basic 01_基础 channel
package basic

import (
	"fmt"
	"time"
)

// 无缓冲通道
func unbufferedChannel() {
	// 创建一个无缓冲通道
	ch := make(chan int)

	// 启动一个 goroutine 来发送数据
	go func() {
		fmt.Println("sending...")
		ch <- 20
		fmt.Println("sent")
	}()

	fmt.Println("receive...")
	value := <-ch
	fmt.Println("receive data:", value)

	// 执行结果
	// receive...
	// sending...
	// sent
	// receive data: 20

	// 示例中，发送方必须等到接收方准备好（goroutine 中 <-ch 操作）才能发送数据
	// 没有缓冲区的通道确保了发送方和接收方的同步
}

// 有缓冲通道
func bufferedChannel() {
	ch := make(chan int, 2) // 创建一个缓冲区大小为 2 的通道
	// 启动一个 goroutine 来发送数据
	go func() {
		fmt.Println("sending 1st data...")
		ch <- 20 // 不会阻塞，数据会存入缓冲区
		fmt.Println("sending 2nd data...")
		ch <- 21 // 不会阻塞，数据会存入缓冲区
		fmt.Println("sending 3rd data...")
		ch <- 22 // 阻塞，直到接收方接收数据
	}()

	fmt.Println("receiving 1st data...")
	value1 := <-ch // 阻塞，直到接收到第一个数据
	fmt.Println("receiving 1st data:", value1)

	fmt.Println("receiving 2nd data...")
	value2 := <-ch // 阻塞，直到接收到第二个数据
	fmt.Println("receiving 2nd data:", value2)

	fmt.Println("receiving 3rd data...")
	value3 := <-ch // 阻塞，直到接收到第三个数据
	fmt.Println("receiving 3rd data:", value3)

	// receiving 1st data...
	// sending 1st data...
	// sending 2nd data...
	// sending 3rd data...
	// receiving 1st data: 20
	// receiving 2nd data...
	// receiving 2nd data: 21
	// receiving 3rd data...
	// receiving 3rd data: 22
}

// 模拟工作池中的每个工作 goroutine
func worker(id int, ch chan<- bool) {
	fmt.Println("worker", id, "is doing work...")
	ch <- true // 发送完成信号
}

// 处理任务并返回结果
func workerPool(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("worker %d is processing job %d\n", id, job)
		results <- job * 2 // 处理任务并发送结果
	}
}

// 将任务提交到通道
func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}

// 从通道接收并处理任务
func consumer(ch <-chan int) {
	for value := range ch {
		fmt.Println("consumer received:", value)
	}
}

// 工作完成信号
func signalSync() {
	ch := make(chan bool, 3)

	// 启动多个 goroutine
	for i := 1; i <= 3; i++ {
		go worker(i, ch)
	}

	// 等待所有 goroutine 完成
	for i := 1; i <= 3; i++ {
		<-ch // 阻塞，直到接收到所有 goroutine 的完成信号
	}
	fmt.Println("all workers are done.")
}

// 任务处理
func taskPool() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// 启动 3 个工作 goroutine
	for w := 1; w <= 3; w++ {
		go workerPool(w, jobs, results)
	}

	// 提交任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs) // 关闭 jobs 通道，表示没有更多的任务

	// 收集结果
	for a := 1; a <= 5; a++ {
		fmt.Println("result:", <-results)
	}
}

// 选择哪个通道首先收到数据
func selectExample() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 2
	}()

	select {
	case res1 := <-ch1:
		fmt.Println("received from ch1:", res1)
	case res2 := <-ch2:
		fmt.Println("received from ch2:", res2)
	}
}

// 通道与管道模式，生产者消费者模型
func pipeModel() {
	ch3 := make(chan int)
	go producer(ch3)
	consumer(ch3)
}

func channel() {
	// 无缓冲通道
	unbufferedChannel()

	// 有缓冲通道
	bufferedChannel()

	// 信号同步
	signalSync()

	// 工作池
	taskPool()

	// select 复用
	selectExample()

	// 通道与管道模式
	pipeModel()
}

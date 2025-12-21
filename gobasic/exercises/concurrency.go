// Package exercises 03_并发编程练习题
package exercises

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

// 计数器
func counter(val *int, m *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	m.Lock()
	for range 1000 {
		*val++
	}
	m.Unlock()
}

func exercises1() {
	var (
		wg     sync.WaitGroup
		m      sync.Mutex
		result int
	)

	for range 10 {
		wg.Add(1)
		go counter(&result, &m, &wg)
	}

	wg.Wait()

	fmt.Println("result:", result)
}

// 生产者/消费者
const numTimes int = 10

func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range numTimes {
		ch <- i
		fmt.Println("producer:", i)
		time.Sleep(time.Millisecond * 100) // 模拟生产延迟
	}
	close(ch)
}

func cunsumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range ch {
		fmt.Println("cunsumer:", item)
		time.Sleep(time.Millisecond * 200) // 模拟消费延迟
	}
}

func exercises2() {
	var wg sync.WaitGroup
	ch := make(chan int, 5)

	wg.Add(2)

	// 生产者
	go producer(ch, &wg)

	// 消费者
	go cunsumer(ch, &wg)

	// 等待所有 goroutine 完成
	wg.Wait()

	fmt.Println("completed")
}

// 并发求和
func sumPart(arr []int, start, end int, result *int64, wg *sync.WaitGroup) {
	defer wg.Done()

	var partSum int64
	for i := start; i < end; i++ {
		partSum += int64(arr[i])
	}

	// 线程安全的更新result，当然 sync.Mutex 也可以
	atomic.AddInt64(result, partSum)
}

func exercises3() {
	// 创建一个数组
	arr := make([]int, 1000000)
	for i := range len(arr) {
		arr[i] = i + 1
	}

	// 扩展下，使用核数
	numCPUs := runtime.NumCPU()
	fmt.Printf("number of CPU cores: %d\n", numCPUs)

	// goroutine数量
	// numGoroutines := 10
	numGoroutines := numCPUs
	chunkSize := len(arr) / numGoroutines
	var wg sync.WaitGroup
	var result int64

	for i := range numGoroutines {
		start := i * chunkSize
		end := (i + 1) * chunkSize
		if i == numGoroutines-1 {
			// 确保最后一部分包含数组的所有剩余元素
			end = len(arr)
		}

		wg.Add(1)
		go sumPart(arr, start, end, &result, &wg)
	}

	// 等待所有 goroutine 完成
	wg.Wait()

	// 输出最终结果
	fmt.Println("total sum:", result)
}

// 读取并发锁
type sharedResource struct {
	counter int
	mu      sync.RWMutex
}

func (sr *sharedResource) readCounter(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	// 读锁
	sr.mu.RLock()
	defer sr.mu.RUnlock()

	// 模拟读取的延迟
	fmt.Printf("reader %d: counter value is %d\n", id, sr.counter)
	time.Sleep(time.Millisecond * 100)
}

func (sr *sharedResource) writerCounter(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	// 写锁
	sr.mu.Lock()
	defer sr.mu.Unlock()

	// 模拟写入的延迟
	sr.counter++
	fmt.Printf("writer %d: counter value updated to %d\n", id, sr.counter)
	time.Sleep(time.Millisecond * 200)
}

func exercises4() {
	var wg sync.WaitGroup
	sr := &sharedResource{counter: 0}

	// 启动多个读取 goroutine
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go sr.readCounter(i, &wg)
	}

	// 写 goroutine
	wg.Add(1)
	go sr.writerCounter(1, &wg)

	// 再启动一些读取操作
	for i := 6; i <= 10; i++ {
		wg.Add(1)
		go sr.readCounter(i, &wg)
	}

	// 等待所有 goroutine 完成
	wg.Wait()

	fmt.Println("final counter value:", sr.counter)
}

// 并发文件下载
type downloadResult struct {
	URL    string
	Status bool
	Err    error
}

func downloadFile(url string, ch chan<- downloadResult, wg *sync.WaitGroup) {
	defer wg.Done()

	// 下载
	resp, err := http.Get(url)
	if err != nil {
		ch <- downloadResult{URL: url, Err: err}
		return
	}
	defer resp.Body.Close()

	// 创建文件
	fileName := url[strings.LastIndex(url, "/")+1:]
	file, err := os.Create(fileName)
	if err != nil {
		ch <- downloadResult{URL: url, Status: false, Err: err}
		return
	}
	defer file.Close()

	// 写入文件
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		ch <- downloadResult{URL: url, Status: false, Err: err}
		return
	}

	// 传递成功信息
	ch <- downloadResult{URL: url, Status: true, Err: nil}
}

func exercises5() {
	// 文件列表
	urls := []string{
		"https://golang.google.cn/images/logos/google.svg",
		"https://golang.google.cn/images/gophers/ladder.svg",
	}

	var wg sync.WaitGroup
	ch := make(chan downloadResult, len(urls))

	// 启动并发下载
	for _, url := range urls {
		wg.Add(1)
		go downloadFile(url, ch, &wg)
	}

	// 等待所有 goroutine 完成
	go func() {
		wg.Wait()
		close(ch) // 下载完成后关闭通道
	}()

	// 打印下载结果
	for result := range ch {
		if !result.Status {
			fmt.Printf("error downloading %s: %v\n", result.URL, result.Err)
		} else {
			fmt.Printf("successfully downloaded %s\n", result.URL)
		}
	}
}

// 并发排序
func mergeSortedSlices(slice1, slice2 []int) []int {
	result := make([]int, 0, len(slice1)+len(slice2))

	i, j := 0, 0
	for i < len(slice1) && j < len(slice2) {
		if (slice1[i]) < slice2[j] {
			result = append(result, slice1[i])
			i++
		} else {
			result = append(result, slice2[j])
			j++
		}
	}

	// append any remaining elements
	result = append(result, slice1[i:]...)
	result = append(result, slice2[j:]...)

	return result
}

func sortSlice(slice []int, wg *sync.WaitGroup) {
	defer wg.Done()
	sort.Ints(slice)
}

func exercises6() {
	// rand.Seed(time.Now().UnixNano()) // go1.20+ 弃用
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]int, 1000000)
	for i := range arr {
		// arr[i] = rand.Intn(1000000)
		arr[i] = randGen.Intn(1000000)
	}

	numGoroutines := runtime.NumCPU()
	chunkSize := len(arr) / numGoroutines
	var wg sync.WaitGroup
	slices := make([][]int, numGoroutines)

	// 将数组分割成多个子数组
	for i := range numGoroutines {
		start := i * chunkSize
		end := (i + 1) * chunkSize
		if i == numGoroutines-1 {
			end = len(arr)
		}
		slices[i] = arr[start:end]
	}

	// 使用 goroutine 对每个子数组进行排序
	for i := range numGoroutines {
		wg.Add(1)
		go sortSlice(slices[i], &wg)
	}

	// 等待所有排序完成
	wg.Wait()

	// 合并所有已排序的子数组
	sortedArray := slices[0]
	for i := 1; i < numGoroutines; i++ {
		sortedArray = mergeSortedSlices(sortedArray, slices[i])
	}

	// 打印部分结果以确认
	fmt.Println("first 20 sorted elements:", sortedArray[:20])
}

// 并发实现 Web 服务器
func handler(w http.ResponseWriter, r *http.Request, wg *sync.WaitGroup) {
	defer wg.Done()
	// 模拟耗时处理
	time.Sleep(2 * time.Second)
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

func exercises7() {
	var wg sync.WaitGroup

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		wg.Add(1)
		handler(w, r, &wg)
	})

	server := &http.Server{
		Addr:    ":1221",
		Handler: mux,
	}

	// 启动服务器（在 goroutine 中）
	go func() {
		fmt.Println("server is running at http://localhost:1221/")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("error starting server:", err)
		}
	}()

	// 等待中断信号（Ctrl+C）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("\nshutting down server...")

	// 给服务器最多 10 秒时间处理完现有请求
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 优雅关闭服务器
	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("server forced to shutdown:", err)
	} else {
		fmt.Println("server stopped accepting new requests.")
	}

	// 等待所有正在处理的请求完成
	fmt.Println("waiting for all active requests to finish...")
	wg.Wait()
	fmt.Println("all requests completed. Goodbye!")
}

// 读写并发模拟
func exercises8() {}

// 限流器实现
func exercises9() {}

// 并发日志记录
func exercises10() {}

// 并发游戏模拟
func exercises11() {}

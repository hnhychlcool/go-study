package main

import (
	"sync"
	"fmt"
	"sync/atomic"
	"runtime"
)

var (
	safeCount int64
	safeWg    sync.WaitGroup
)

/**
*	通过原子函数自增来确保资源的安全访问
 */
func main() {

	fmt.Println("start automatic goroutine...")
	// 一定要指定等待执行完成的线程数量
	safeWg.Add(2)

	go automaticIncreaseCount()
	go automaticIncreaseCount()

	// 一定要调用等待线程执行完成的wait
	safeWg.Wait()

	fmt.Printf("finished goroutine counter: %d", safeCount)
}

func automaticIncreaseCount() {
	// wait group 完成时,在函数退出时,通知main一个消息
	defer safeWg.Done()

	for i := 0; i < 10000; i++ {
		atomic.AddInt64(&safeCount, 1)
		runtime.Gosched()
	}
}

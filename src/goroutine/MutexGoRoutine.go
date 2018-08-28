package main

import (
	"sync"
	"fmt"
	"runtime"
)

var (
	mutexCount int64
	mutexWg    sync.WaitGroup
	mutex      sync.Mutex
)

/**
*	通过原子函数自增来确保资源的安全访问
 */
func main() {

	fmt.Println("start automatic goroutine...")
	// 一定要指定等待执行完成的线程数量
	mutexWg.Add(2)

	go mutexIncreaseCount()
	go mutexIncreaseCount()

	// 一定要调用等待线程执行完成的wait
	mutexWg.Wait()

	fmt.Printf("finished goroutine counter: %d", mutexCount)
}

func mutexIncreaseCount() {
	// wait group 完成时,在函数退出时,通知main一个消息
	defer mutexWg.Done()

	for i := 0; i < 10000; i++ {
		// 设置临界区
		mutex.Lock()
		{
			mutexCount++
			runtime.Gosched()
		}
		// 释放临界区
		mutex.Unlock()
	}
}

package main

import (
	"sync"
	"math/rand"
	"time"
	"fmt"
)

var bufferedWg sync.WaitGroup

const (
	taskLoad       = 10
	goroutineCount = 10
)

// init 初始化包,Go运行时,会在其他代码执行前,执行这个方法
func init() {
	// 初始化随机数种子
	rand.Seed(time.Now().Unix())
}

func main() {
	bufferedWg.Add(goroutineCount)

	// 创建一个有缓冲的通道
	tasks := make(chan string, taskLoad)

	for i := 1; i <= goroutineCount; i++ {
		go work(i, tasks)
	}

	// 增加一组工作
	for post := 1; post < goroutineCount; post++ {
		tasks <- fmt.Sprintf("Task %d posted", post)
	}

	// 当所有工作都处理完时,关闭通道
	// 以便所有的goroutine都能退出
	close(tasks)

	bufferedWg.Wait()
}

func work(work int, tasks chan string) {
	defer bufferedWg.Done()

	for {
		task, ok := <-tasks
		if !ok {
			fmt.Printf("work: %d shut down \n", work)
			return
		}
		fmt.Printf("work: %d started %s \n", work, task)

		// 随机一段时间来模拟工作时间
		sleep := rand.Int63n(1000)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Printf("work: %d finished %s \n", work, task)
	}
}

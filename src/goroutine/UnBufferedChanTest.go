package main

import (
	"sync"
	"fmt"
	"time"
)

var unBufferedChanWg sync.WaitGroup

func main() {

	unBufferedChanWg.Add(1)

	// 创建一个无缓冲通道
	buton := make(chan int)
	// 开始比赛
	go Runner(buton)

	buton <- 1

	unBufferedChanWg.Wait()

}

func Runner(baton chan int) {
	var newRunner int
	runner := <-baton

	if runner != 4 {
		newRunner = runner + 1
		go Runner(baton)
	}

	time.Sleep(100 * time.Millisecond)

	if runner == 4 {
		fmt.Printf("Runner %d finished race \n", runner)
		unBufferedChanWg.Done()
		return
	}

	fmt.Printf("exchange to %d to run \n", newRunner)
	baton <- newRunner

}

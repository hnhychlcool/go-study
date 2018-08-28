package main

import (
	"runtime"
	"fmt"
	"sync"
)

var count int
var raceConditionWg sync.WaitGroup
func main() {
	runtime.GOMAXPROCS(2)
	raceConditionWg.Add(2)

	fmt.Printf("start goroutine...")

	go increaseCount(1)
	go increaseCount(2)

	raceConditionWg.Wait()
	fmt.Printf("finished goroutine, count=%d", count)
}

func increaseCount(param int) {
	defer raceConditionWg.Done()

	for i := 0; i < 10000; i++ {
		runtime.Gosched()
		count++
	}
}

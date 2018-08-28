package main

import (
	"sync"
	"fmt"
	"time"
	"sync/atomic"
)

var (
	flag        int64
	storeLoadWg sync.WaitGroup
)

func main() {
	storeLoadWg.Add(2)

	go storeAndLoad("chl")
	go storeAndLoad("xuhongli")

	time.Sleep(1 * time.Second)
	println("shutdown now...")
	atomic.StoreInt64(&flag, 1)

	storeLoadWg.Wait()
	println("store and load safe goroutine finished...")
}
func storeAndLoad(name string) {
	defer storeLoadWg.Done()
	for {
		fmt.Printf("doing %s work \n", name)
		time.Sleep(250 * time.Millisecond)
		if atomic.LoadInt64(&flag) == 1 {
			fmt.Printf("finish %s work, exit \n", name)
			break
		}
	}

}

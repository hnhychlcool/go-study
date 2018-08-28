//package main
package main
import (
	"sync"
	"fmt"
)

var wg sync.WaitGroup
/*
func main() {

	runtime.GOMAXPROCS(2)

	// 计数器加2,表示要等待2个goroutine
	wg.Add(2)

	fmt.Println("Create goroutine...")

	go primeA("A")
	go primeA("B")

	fmt.Println("Waiting goroutine finish...")
	wg.Wait()
	fmt.Println("all goroutine finished...")

}*/

func primeA(s string) {
	defer wg.Done()

next:
	for outter := 2; outter < 5000; outter++ {
		for inner := outter; inner < outter; inner++ {
			if outter%inner == 0 {
				continue next;
			}
		}
		fmt.Printf("%s %d \n", s, outter)
	}
	fmt.Println("completed", s);
}

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func atomic_count_test() {
	var count int64 = 0
	var wg sync.WaitGroup

	increment := func() {
		for j := 0; j < 1000; j++ {
			atomic.AddInt64(&count, 1)
		}
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}
	wg.Wait()
	fmt.Println("Final Count:", count)
}

func main() {
	atomic_count_test()
}

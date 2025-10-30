package main

import (
	"fmt"
	"sync"
)

func count_test() {
	var mu sync.Mutex
	count := 0
	var wg sync.WaitGroup

	increment := func() {
		for j := 0; j < 1000; j++ {
			mu.Lock()
			count++
			mu.Unlock()
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
	count_test()
}

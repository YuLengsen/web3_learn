package main

import (
	"fmt"
	"sync"
	"time"
)

type Task func()

func runTasksConcurrently(tasks []Task) {
	var wg sync.WaitGroup
	for _, task := range tasks {
		wg.Add(1)
		go func(t Task) {
			defer wg.Done()
			start := time.Now()
			t()
			fmt.Printf("Task completed in %v\n", time.Since(start))
		}(task)
	}
	wg.Wait()
}

func main() {
	tasks := []Task{
		func() {
			for i := 0; i < 5; i++ {
				fmt.Println("Task 1 - Count:", i)
				time.Sleep(100 * time.Millisecond)
			}
		},
		func() {
			for i := 0; i < 5; i++ {
				fmt.Println("Task 2 - Count:", i)
				time.Sleep(150 * time.Millisecond)
			}
		},
		func() {
			for i := 0; i < 5; i++ {
				fmt.Println("Task 3 - Count:", i)
				time.Sleep(200 * time.Millisecond)
			}
		},
	}

	runTasksConcurrently(tasks)
}

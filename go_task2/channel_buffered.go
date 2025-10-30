package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
		fmt.Println("Produced:", i)
	}
	close(ch)
}

func consumer(ch chan int) {
	for val := range ch {
		fmt.Println("Consumed:", val)
		time.Sleep(50 * time.Millisecond) // Simulate processing time
	}
}

func main() {
	bufferedChan := make(chan int, 10) // Buffered channel with capacity of 10

	go producer(bufferedChan)
	consumer(bufferedChan)
}

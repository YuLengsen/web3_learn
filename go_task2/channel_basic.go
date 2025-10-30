package main

import "fmt"

func producter(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("Produced:", i)
	}
	close(ch)
}

func consumer(ch chan int) {
	for val := range ch {
		fmt.Println("Consumed:", val)
	}
}

func main() {
	bufferedChan := make(chan int)

	go producter(bufferedChan)
	consumer(bufferedChan)
}

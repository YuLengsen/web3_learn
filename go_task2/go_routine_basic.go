package main

import (
	"fmt"
	"time"
)

func printOddNumbers() {
	for i := 1; i < 10; i += 2 {
		fmt.Println("Odd Number:", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printEvenNumbers() {
	for i := 0; i < 10; i += 2 {
		fmt.Println("Even Number:", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go printOddNumbers()
	go printEvenNumbers()

	// Wait for goroutines to finish
	time.Sleep(2 * time.Second)
}

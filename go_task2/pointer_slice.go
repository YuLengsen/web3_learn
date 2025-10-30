package main

import "fmt"

func doubleSlice(slicePtr *[]int) {
	slice := *slicePtr
	for i := range slice {
		slice[i] *= 2
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	doubleSlice(&numbers)
	fmt.Println("Doubled slice:", numbers)
}

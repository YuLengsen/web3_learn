package main

import "fmt"

func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	single := singleNumber(nums)
	fmt.Println("只出现一次的数字是:", single)
}

package main

import "fmt"

func addTen(num *int) {
	*num += 10
}

func main() {
	number := 5
	addTen(&number)
	fmt.Println("Number after adding ten:", number)
}

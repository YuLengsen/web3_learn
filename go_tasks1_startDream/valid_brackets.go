package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else if char == ')' || char == '}' || char == ']' {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	testStr := "{[()]}"
	if isValid(testStr) {
		fmt.Println(testStr, "是有效的括号字符串")
	} else {
		fmt.Println(testStr, "不是有效的括号字符串")
	}
}

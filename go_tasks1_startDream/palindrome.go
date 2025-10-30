package main

/*
*
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

回文整数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

例如，121 是回文，而 123 不是。
*/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	original := x
	reversed := 0

	for x != 0 {
		digit := x % 10
		reversed = reversed*10 + digit
		x /= 10
	}

	return original == reversed
}

func main() {
	x := 121
	if isPalindrome(x) {
		println("是回文整数")
	} else {
		println("不是回文整数")
	}
}

package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for _, str := range strs[1:] {
		prefix = prefix[:min(len(str), len(prefix))]
		for len(prefix) > 0 && str[:min(len(prefix), len(str))] != prefix[:min(len(prefix), len(str))] {
			prefix = prefix[:min(len(str), len(prefix)-1)]
		}
		if len(prefix) == 0 {
			return ""
		}
	}
	return prefix
}

func main() {
	strs := []string{"ab", "a"}
	result := longestCommonPrefix(strs)
	fmt.Println("Longest Common Prefix:", result)
}

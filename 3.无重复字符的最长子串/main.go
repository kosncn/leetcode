package main

import "fmt"

func main() {
	// 示例 1
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s)) // Output: 3

	// 示例 2
	s = "bbbbb"
	fmt.Println(lengthOfLongestSubstring(s)) // Output: 1

	// 示例 3
	s = "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s)) // Output: 3

	// 示例 4
	s = " "
	fmt.Println(lengthOfLongestSubstring(s)) // Output: 1
}

func lengthOfLongestSubstring(s string) int {
	// 方法一: 暴力
	count := 0
	for i := 0; i < len(s); i++ {
		char := make(map[byte]bool)
		for j := i; j < len(s); j++ {
			if _, ok := char[s[j]]; ok {
				break
			}
			char[s[j]] = true
		}
		if count < len(char) {
			count = len(char)
		}
	}
	return count
}

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
	// 解法一: 暴力
	//result := 0
	//for i := 0; i < len(s); i++ {
	//	char := make(map[byte]bool)
	//	for j := i; j < len(s); j++ {
	//		if _, ok := char[s[j]]; ok {
	//			break
	//		}
	//		char[s[j]] = true
	//	}
	//	if result < len(char) {
	//		result = len(char)
	//	}
	//}
	//return result

	// 解法二: 滑动窗口
	result := 0
	left, right := 0, 0
	window := make(map[byte]int)
	for right < len(s) {
		c1 := s[right]
		window[c1]++
		right++
		for window[c1] > 1 {
			c2 := s[left]
			window[c2]--
			left++
		}
		if right-left > result {
			result = right - left
		}
	}
	return result
}

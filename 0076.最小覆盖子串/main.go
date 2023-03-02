package main

import (
	"fmt"
	"math"
)

func main() {
	// 示例 1
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s, t)) // Output: "BANC"

	// 示例 2
	s = "a"
	t = "a"
	fmt.Println(minWindow(s, t)) // Output: "a"

	// 示例 3
	s = "a"
	t = "aa"
	fmt.Println(minWindow(s, t)) // Output: ""
}

func minWindow(s string, t string) string {
	// 解法一：滑动窗口
	window, need := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	left, right := 0, 0
	start, end := 0, math.MaxInt
	valid := 0
	for right < len(s) {
		c1 := s[right]
		right++
		if _, ok := need[c1]; ok {
			window[c1]++
			if window[c1] == need[c1] {
				valid++
			}
		}
		for valid == len(need) {
			if right-left < end-start {
				start, end = left, right
			}
			c2 := s[left]
			left++
			if _, ok := need[c2]; ok {
				if window[c2] == need[c2] {
					valid--
				}
				window[c2]--
			}
		}
	}
	if end == math.MaxInt {
		return ""
	}
	return s[start:end]
}

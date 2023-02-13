package main

import "fmt"

func main() {
	// 示例 1
	s := "abc"
	fmt.Println(countSubstrings(s)) // Output: 3

	// 示例 2
	s = "aaa"
	fmt.Println(countSubstrings(s)) // Output: 6

	// 示例 3
	s = "a"
	fmt.Println(countSubstrings(s)) // Output: 1
}

func countSubstrings(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		result += countPalindrome(s, i, i)
	}
	for i := 0; i < len(s)-1; i++ {
		result += countPalindrome(s, i, i+1)
	}
	return result
}

func countPalindrome(s string, left, right int) int {
	count := 0
	for left > -1 && right < len(s) && s[left] == s[right] {
		count += 1
		left, right = left-1, right+1
	}
	return count
}

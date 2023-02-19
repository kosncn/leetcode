package main

import "fmt"

func main() {
	// 示例 1
	s := "babad"
	fmt.Println(longestPalindrome(s)) // Output: bab 或 aba

	// 示例 2
	s = "cbbd"
	fmt.Println(longestPalindrome(s)) // Output: bb
}

func longestPalindrome(s string) string {
	// 解法一: 双指针
	//result := ""
	//length := len(s)
	//for i := 0; i < length; i++ {
	//	s1 := helper(s, i, i)
	//	if len(s1) > len(result) {
	//		result = s1
	//	}
	//	s2 := helper(s, i, i+1)
	//	if len(s2) > len(result) {
	//		result = s2
	//	}
	//}
	//return result

	// 解法二: 马拉车
	str := "^"
	m := len(s)
	for i := 0; i < m; i++ {
		str += "#" + string(s[i])
		if i == m-1 {
			str += "#"
		}
	}
	str += "$"

	n := len(str)
	middle, right := 0, 0
	radius := make([]int, n)
	for i := 1; i < n-1; i++ {
		j := 2*middle - i
		if i < right {
			radius[i] = min(radius[j], right-i)
		}
		for str[i+radius[i]+1] == str[i-radius[i]-1] {
			radius[i]++
		}
		if right < i+radius[i] {
			middle = i
			right = i + radius[i]
		}
	}

	index, length := 0, 0
	for i := 1; i < n-1; i++ {
		if length < radius[i] {
			index = i
			length = radius[i]
		}
	}
	index = (index - length) / 2

	return s[index : index+length]
}

func helper(s string, left, right int) string {
	for left > -1 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

package main

import "fmt"

func main() {
	// 示例 1
	s := "(()"
	fmt.Println(longestValidParentheses(s)) // Output: 2

	// 示例 2
	s = ")()())"
	fmt.Println(longestValidParentheses(s)) // Output: 4

	// 示例 3
	s = ""
	fmt.Println(longestValidParentheses(s)) // Output: 0
}

func longestValidParentheses(s string) int {
	// 解法一：暴力（超时）
	//count := 0
	//for i := 0; i < len(s); i++ {
	//	for j := len(s) - 1; j > i; j-- {
	//		if isValidParentheses(s[i:j+1]) && j-i+1 > count {
	//			count = j - i + 1
	//			break
	//		}
	//	}
	//}
	//return count

	// 解法二：动态规划
	count := 0
	dp := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if s[i] == '(' {
			continue
		}
		if s[i-1] == '(' {
			dp[i] = 2
			if i-2 >= 0 {
				dp[i] += dp[i-2]
			}
		} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
			dp[i] = dp[i-1] + 2
			if i-dp[i-1]-2 >= 0 {
				dp[i] += dp[i-dp[i-1]-2]
			}
		}
		count = max(count, dp[i])
	}
	return count
}

func isValidParentheses(s string) bool {
	var stack []byte
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, s[i])
		} else if len(stack) > 0 && stack[len(stack)-1] == '(' {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

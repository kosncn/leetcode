package main

import "fmt"

var result []string

func main() {
	// 示例 1
	n := 3
	fmt.Println(generateParenthesis(n)) // Output: ["((()))", "(()())", "(())()", "()(())", "()()()"]

	// 示例 2
	n = 1
	fmt.Println(generateParenthesis(n)) // Output: ["()"]
}

func generateParenthesis(n int) []string {
	result = []string{}
	backtrack(n, n, "")
	return result
}

func backtrack(left, right int, current string) {
	if left < 0 || right < 0 || left > right {
		return
	}
	if left == 0 && right == 0 {
		result = append(result, current)
		return
	}

	current += "("
	backtrack(left-1, right, current)
	current = current[:len(current)-1]

	current += ")"
	backtrack(left, right-1, current)
	current = current[:len(current)-1]
}

package main

import "fmt"

func main() {
	// 示例 1
	s := "()"
	fmt.Println(isValid(s)) // Output: true

	// 示例 2
	s = "()[]{}"
	fmt.Println(isValid(s)) // Output: true

	// 示例 3
	s = "(]"
	fmt.Println(isValid(s)) // Output: false
}

func isValid(s string) bool {
	// 解法一：栈
	var stack []byte
	var parenthesis = map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && s[i] == parenthesis[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

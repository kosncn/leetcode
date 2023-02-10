package main

import "fmt"

var result []string

var phoneMap = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func main() {
	// 示例 1
	digits := "23"
	fmt.Println(letterCombinations(digits)) // Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"]

	// 示例 2
	digits = ""
	fmt.Println(letterCombinations(digits)) // Output: []

	// 示例 3
	digits = "2"
	fmt.Println(letterCombinations(digits)) // Output: ["a", "b", "c"]
}

func letterCombinations(digits string) []string {
	result = []string{}
	if len(digits) <= 0 {
		return result
	}
	backtrack(digits, 0, "")
	return result
}

func backtrack(digits string, index int, combination string) {
	if len(combination) == len(digits) {
		result = append(result, combination)
		return
	}
	for i := index; i < len(digits); i++ {
		for _, c := range phoneMap[digits[i]] {
			combination += c
			backtrack(digits, i+1, combination)
			combination = combination[:len(combination)-1]
		}
	}
}

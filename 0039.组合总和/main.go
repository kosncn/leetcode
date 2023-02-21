package main

import (
	"fmt"
	"sort"
)

func main() {
	// 示例 1
	candidates := []int{2, 3, 6, 7}
	target := 7
	fmt.Println(combinationSum(candidates, target)) // Output: [[2, 2, 3], [7]]

	// 示例 2
	candidates = []int{2, 3, 5}
	target = 8
	fmt.Println(combinationSum(candidates, target)) // Output: [[2, 2, 2, 2], [2, 3, 3], [3, 5]]

	// 示例 3
	candidates = []int{2}
	target = 1
	fmt.Println(combinationSum(candidates, target)) // Output: []
}

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	sort.Ints(candidates)
	backtrack(candidates, target, 0, []int{}, &result)
	return result
}

func backtrack(candidates []int, target int, index int, numbers []int, result *[][]int) {
	if target == 0 {
		*result = append(*result, append([]int{}, numbers...))
		return
	}
	if target < 0 {
		return
	}
	for i := index; i < len(candidates); i++ {
		numbers = append(numbers, candidates[i])
		backtrack(candidates, target-candidates[i], i, numbers, result)
		numbers = numbers[:len(numbers)-1]
	}
}

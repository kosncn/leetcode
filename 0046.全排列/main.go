package main

import "fmt"

func main() {
	// 示例 1
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums)) // Output: [[1, 2, 3], [1, 3, 2], [2, 1, 3], [2, 3, 1], [3, 1, 2], [3, 2, 1]]

	// 示例 2
	nums = []int{0, 1}
	fmt.Println(permute(nums)) // Output: [[0, 1], [1, 0]]]

	// 示例 3
	nums = []int{1}
	fmt.Println(permute(nums)) // Output: [[1]]
}

func permute(nums []int) [][]int {
	// 解法一：回溯
	var result [][]int
	var used = make(map[int]bool)
	backtrack(nums, []int{}, &used, &result)
	return result
}

func backtrack(nums []int, array []int, used *map[int]bool, result *[][]int) {
	if len(array) == len(nums) {
		*result = append(*result, append([]int{}, array...))
		return
	}
	for _, v := range nums {
		if (*used)[v] {
			continue
		}
		(*used)[v] = true
		array = append(array, v)
		backtrack(nums, array, used, result)
		array = array[:len(array)-1]
		(*used)[v] = false
	}
}

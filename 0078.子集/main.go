package main

import "fmt"

func main() {
	// 示例 1
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums)) // Output: [[], [1], [2], [1, 2], [3], [1, 3], [2, 3], [1, 2, 3]]

	//// 示例 2
	//nums = []int{0}
	//fmt.Println(subsets(nums)) // Output: [[], [0]]
	//
	//// 示例 2
	//nums = []int{9, 0, 3, 5, 7}
	//fmt.Println(subsets(nums)) // Output: [[], [0]]
}

func subsets(nums []int) [][]int {
	// 解法一：迭代
	result := [][]int{{}}
	for i := 0; i < len(nums); i++ {
		length := len(result)
		for j := 0; j < length; j++ {
			temp := make([]int, len(result[j]))
			copy(temp, result[j])
			result = append(result, append(temp, nums[i]))
		}
	}
	return result

	// 解法二：回溯
	//var result [][]int
	//var backtrack func(int, []int)
	//backtrack = func(index int, track []int) {
	//	result = append(result, append([]int{}, track...))
	//	for i := index; i < len(nums); i++ {
	//		track = append(track, nums[i])
	//		backtrack(i+1, track)
	//		track = track[:len(track)-1]
	//	}
	//}
	//backtrack(0, []int{})
	//return result
}

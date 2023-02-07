package main

import (
	"fmt"
	"sort"
)

func main() {
	// 示例 1
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums)) // Output: [[-1, -1, 2], [-1, 0, 1]]

	// 示例 2
	nums = []int{0, 1, 1}
	fmt.Println(threeSum(nums)) // Output: []

	// 示例 3
	nums = []int{0, 0, 0}
	fmt.Println(threeSum(nums)) // Output: [[0, 0, 0]]
}

func threeSum(nums []int) [][]int {
	// 解法一: 排序 + 双指针
	var result [][]int
	length := len(nums)
	sort.Ints(nums)
	for i := 0; i < length-2; i++ {
		if nums[i]+nums[i+1]+nums[i+2] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, length-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left+1] == nums[left] {
					left++
				}
				for left < right && nums[right-1] == nums[right] {
					right--
				}
				left, right = left+1, right-1
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

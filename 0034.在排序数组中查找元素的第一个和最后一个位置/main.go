package main

import "fmt"

func main() {
	// 示例 1
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println(searchRange(nums, target)) // Output: [3, 4]

	// 示例 2
	nums = []int{5, 7, 7, 8, 8, 10}
	target = 6
	fmt.Println(searchRange(nums, target)) // Output: [-1, -1]

	// 示例 3
	nums = []int{}
	target = 0
	fmt.Println(searchRange(nums, target)) // Output: [-1, -1]
}

func searchRange(nums []int, target int) []int {
	left := searchLeftBound(nums, target)
	right := searchRightBound(nums, target)
	return []int{left, right}
}

func searchLeftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := left + (right-left)/2
		if nums[middle] >= target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	if left > len(nums)-1 || nums[left] != target {
		return -1
	}
	return left
}

func searchRightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := left + (right-left)/2
		if nums[middle] <= target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	if right < 0 || nums[right] != target {
		return -1
	}
	return right
}

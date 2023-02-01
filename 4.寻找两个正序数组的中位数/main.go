package main

import (
	"fmt"
)

func main() {
	// 示例 1
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2)) // Output: 2.0

	// 示例 2
	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2)) // Output: 2.5

	// 示例 3
	nums1 = []int{0, 0}
	nums2 = []int{0, 0}
	fmt.Println(findMedianSortedArrays(nums1, nums2)) // Output: 0.0
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 双指针
	p1, p2 := 0, 0
	c1, c2 := len(nums1), len(nums2)
	left, right := -1, -1
	length := c1 + c2
	for i := 0; i < length/2+1; i++ {
		left = right
		if p1 < c1 && (p2 >= c2 || nums1[p1] < nums2[p2]) {
			right = nums1[p1]
			p1++
		} else {
			right = nums2[p2]
			p2++
		}
	}
	if length%2 == 0 {
		return float64(left+right) / 2.0
	}
	return float64(right)
}

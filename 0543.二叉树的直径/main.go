package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// 示例 1
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println(diameterOfBinaryTree(root)) // Output: 3
}

func diameterOfBinaryTree(root *TreeNode) int {
	// 解法一：回溯算法思路
	var diameter int
	var traverse func(*TreeNode) int
	traverse = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := traverse(node.Left)
		right := traverse(node.Right)
		diameter = max(diameter, left+right)
		return 1 + max(left, right)
	}
	traverse(root)
	return diameter
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

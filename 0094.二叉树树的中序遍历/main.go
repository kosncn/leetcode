package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// 示例 1
	root := &TreeNode{ // root: [1, null, 2, 3]
		Val: 1,
		Right: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 3},
		},
	}
	fmt.Println(inorderTraversal(root)) // Output: [1, 3, 2]

	// 示例 2
	root = nil
	fmt.Println(inorderTraversal(root)) // Output: []

	// 示例 3
	root = &TreeNode{Val: 1}
	fmt.Println(inorderTraversal(root)) // Output: [1]
}

func inorderTraversal(root *TreeNode) []int {
	// 解法一：递归
	//var result []int
	//var inorder func(node *TreeNode)
	//inorder = func(node *TreeNode) {
	//	if node == nil {
	//		return
	//	}
	//	inorder(node.Left)
	//	result = append(result, node.Val)
	//	inorder(node.Right)
	//}
	//inorder(root)
	//return result

	// 解法二：迭代
	//var result []int
	//var stack []*TreeNode
	//for root != nil || len(stack) > 0 {
	//	for root != nil {
	//		stack = append(stack, root)
	//		root = root.Left
	//	}
	//	root = stack[len(stack)-1]
	//	stack = stack[:len(stack)-1]
	//	result = append(result, root.Val)
	//	root = root.Right
	//}
	//return result

	// 解法三：Morris 中序遍历
	var result []int
	for root != nil {
		if root.Left != nil {
			node := root.Left
			for node.Right != nil && node.Right != root {
				node = node.Right
			}
			if node.Right == nil {
				node.Right = root
				root = root.Left
			} else {
				result = append(result, root.Val)
				node.Right = nil
				root = root.Right
			}
		} else {
			result = append(result, root.Val)
			root = root.Right
		}
	}
	return result
}

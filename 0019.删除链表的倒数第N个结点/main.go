package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// 示例 1
	head := &ListNode{ // head: [1, 2, 3, 4, 5]
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: &ListNode{Val: 5},
				},
			},
		},
	}
	n := 2
	fmt.Println(printListNode(removeNthFromEnd(head, n))) // Output: [1, 2, 3, 5]

	// 示例 2
	head = &ListNode{Val: 1} // head: [1]
	n = 1
	fmt.Println(printListNode(removeNthFromEnd(head, n))) // Output: []

	// 示例 3
	head = &ListNode{ // head: [1, 2]
		Val:  1,
		Next: &ListNode{Val: 2},
	}
	n = 1
	fmt.Println(printListNode(removeNthFromEnd(head, n))) // Output: [1]
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 解法一：快慢指针
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	slow, fast := dummy, dummy
	for i := 0; i < n+1; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow, fast = slow.Next, fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

func printListNode(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

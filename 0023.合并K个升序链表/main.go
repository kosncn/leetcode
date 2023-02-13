package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// 示例 1
	lists := []*ListNode{ // lists: [[1, 4, 5], [1, 3, 4], [2, 6]]
		{
			Val: 1,
			Next: &ListNode{
				Val:  4,
				Next: &ListNode{Val: 5},
			},
		},
		{
			Val: 1,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{Val: 4},
			},
		},
		{
			Val:  2,
			Next: &ListNode{Val: 6},
		},
	}
	fmt.Println(printListNode(mergeKLists(lists))) // Output: [1, 1, 2, 3, 4, 4, 5, 6]

	// 示例 2
	lists = []*ListNode{}                          // lists: []
	fmt.Println(printListNode(mergeKLists(lists))) // Output: []

	// 示例 3
	lists = []*ListNode{nil}                       // lists: [[]]
	fmt.Println(printListNode(mergeKLists(lists))) // Output: []
}

func mergeKLists(lists []*ListNode) *ListNode {
	// 解法一：顺序合并
	var result *ListNode
	for i := 0; i < len(lists); i++ {
		result = mergeTowLists(result, lists[i])
	}
	return result

	// 解法二：优先级队列
}

func mergeTowLists(list1, list2 *ListNode) *ListNode {
	dummy := new(ListNode)
	p, p1, p2 := dummy, list1, list2
	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		p = p.Next
	}
	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}
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

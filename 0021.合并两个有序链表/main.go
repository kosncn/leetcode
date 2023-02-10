package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// 示例 1
	l1 := &ListNode{ // l1: [1, 2, 4]
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{Val: 4},
		},
	}
	l2 := &ListNode{ // l2: [1, 3, 4]
		Val: 1,
		Next: &ListNode{
			Val:  3,
			Next: &ListNode{Val: 4},
		},
	}
	fmt.Println(printListNode(mergeTwoLists(l1, l2))) // Output: [1, 1, 2, 3, 4, 4]

	// 示例 2
	l1 = nil                                          // l1: []
	l2 = nil                                          // l2: []
	fmt.Println(printListNode(mergeTwoLists(l1, l2))) // Output: []

	// 示例 3
	l1 = nil                                          // l1: []
	l2 = &ListNode{Val: 0}                            // l2: [0]
	fmt.Println(printListNode(mergeTwoLists(l1, l2))) // Output: [0]

	// 示例 4
	l1 = &ListNode{Val: 5} // l1: [5]
	l2 = &ListNode{ // l2: [1, 2, 4]
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{Val: 4},
		},
	}
	fmt.Println(printListNode(mergeTwoLists(l1, l2))) // Output: [1, 2, 4, 5]
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	//// 解法一：递归
	//p1, p2 := list1, list2
	//if p1 == nil {
	//	return p2
	//} else if p2 == nil {
	//	return p1
	//} else if p1.Val <= p2.Val {
	//	p1.Next = mergeTwoLists(p1.Next, p2)
	//	return p1
	//} else {
	//	p2.Next = mergeTwoLists(p1, p2.Next)
	//	return p2
	//}

	// 解法二：双指针
	var dummy = new(ListNode)
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

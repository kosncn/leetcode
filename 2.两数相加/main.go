package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	p1, p2 := l1, l2
	bit := 0
	for p1 != nil || p2 != nil {
		sum := p1.Val + p2.Val

	}
}

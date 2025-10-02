package p2807

import . "leetcode/stucture"

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	prev := head
	cur := prev.Next
	for cur != nil {
		newVal := greatestCommonDivisor(prev.Val, cur.Val)
		newNode := &ListNode{Val: newVal, Next: cur}
		prev.Next = newNode
		prev = cur
		cur = cur.Next
	}
	return head
}

func greatestCommonDivisor(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func greatestCommonDivisor0(a, b int) int {
	if a < b {
		a, b = b, a
	}
	o := a % b
	for o > 0 {
		a, b = b, o
		o = a % b
	}
	return b
}

func run(head []int) []int {
	return ListToArr(insertGreatestCommonDivisors(ArrToList(head)))
}

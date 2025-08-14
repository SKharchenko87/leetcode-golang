package main

import . "leetcode/stucture"

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var newListNode, prev, res *ListNode
	var p1, p2 *ListNode
	p1 = list1
	p2 = list2

	for !(p1 == nil && p2 == nil) {
		if p1 != nil && p2 != nil {
			if p1.Val < p2.Val {
				newListNode = &ListNode{Val: p1.Val, Next: nil}
				if prev != nil {
					prev.Next = newListNode
				}
				p1 = p1.Next
			} else {
				newListNode = &ListNode{Val: p2.Val, Next: nil}
				if prev != nil {
					prev.Next = newListNode
				}
				p2 = p2.Next
			}
			if res == nil {
				res = newListNode
			}
			prev = newListNode
		}
		if p1 == nil {
			for p2 != nil {
				newListNode = &ListNode{Val: p2.Val, Next: nil}
				if prev != nil {
					prev.Next = newListNode
				}
				p2 = p2.Next
				if res == nil {
					res = newListNode
				}
				prev = newListNode
			}
		}

		if p2 == nil {
			for p1 != nil {
				newListNode = &ListNode{Val: p1.Val, Next: nil}
				if prev != nil {
					prev.Next = newListNode
				}
				p1 = p1.Next
				if res == nil {
					res = newListNode
				}
				prev = newListNode
			}
		}

	}

	return res
}

func main() {

	m1 := []int{1, 2, 4}
	m2 := []int{1, 3, 4}

	ln1 := SliceToListNode(m1)
	ln2 := SliceToListNode(m2)

	//PrintListNode(ln1)
	//PrintListNode(ln2)
	PrintListNode(mergeTwoLists(ln1, ln2))
}

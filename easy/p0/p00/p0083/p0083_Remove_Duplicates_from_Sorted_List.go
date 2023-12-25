package main

import . "leetcode/stucture"

func deleteDuplicates(head *ListNode) *ListNode {
	p := head
	if p == nil {
		return nil
	}
	cur := head.Val
	result := ListNode{Val: cur, Next: nil}
	new_head := &result
	for {
		for p != nil && p.Val == cur {
			p = p.Next
		}
		if p == nil {
			return &result
		}
		cur = p.Val
		ln := ListNode{Val: cur, Next: nil}
		new_head.Next = &ln
		new_head = &ln
	}
}

func main() {
	m1 := []int{1, 1, 2, 2, 3, 3, 3}
	ln1 := SliceToListNode(m1)
	PrintListNode(ln1)
	x := deleteDuplicates(ln1)
	PrintListNode(x)

}

package main

import . "leetcode/stucture"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var curln, new ListNode
	var ln *ListNode
	a1 := 0
	a2 := 0
	buf := 0
	ln = &ListNode{}
	new = *ln
	curln = *ln
	for !(l1 == nil && l2 == nil) {
		if l1 == nil {
			a1 = 0
		} else {
			a1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			a2 = 0
		} else {
			a2 = l2.Val
			l2 = l2.Next
		}
		cur := a1 + a2 + buf
		buf = buf / 10
		curln.Val = cur % 10
		curln.Next = &new
		new = ListNode{}
		curln = new

	}
	return ln
}

func main() {
	l1 := ListNode{Val: 3, Next: nil}
	l1 = ListNode{Val: 4, Next: &l1}
	l1 = ListNode{Val: 2, Next: &l1}

	l2 := ListNode{Val: 4, Next: nil}
	l2 = ListNode{Val: 6, Next: &l2}
	l2 = ListNode{Val: 5, Next: &l2}

	nums1 := []int{2, 4, 3}
	var ln1 ListNode
	ln1 = *(ln1.ArrayToListNode(&nums1))
	ln1.Print()

	nums2 := []int{5, 6, 4}
	var ln2 ListNode
	ln2 = *(ln2.ArrayToListNode(&nums2))
	ln2.Print()

	var ln3 ListNode
	ln3 = *addTwoNumbers(&ln1, &ln2)
	ln3.Print()

	//for {
	//	fmt.Println(ln1.Val)
	//	if ln1.Next == nil {
	//		break
	//	}
	//	ln1 = *ln1.Next
	//}
	//fmt.Println(ln1)
}

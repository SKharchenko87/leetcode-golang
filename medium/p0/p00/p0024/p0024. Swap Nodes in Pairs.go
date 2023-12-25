package main

import . "leetcode/stucture"

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var headPoint *ListNode
	headPoint = head
	for head != nil {
		a := head
		if a.Next != nil {
			b := a.Next
			b.Val, a.Val = a.Val, b.Val
			head = b.Next
		} else {
			break
		}

	}
	return headPoint
}

func main() {
	nums1 := []int{1, 2, 4, 3}
	var ln1 ListNode
	ln1 = *(ln1.ArrayToListNode(&nums1))
	ln1.Print()
	PrintListNode(swapPairs(&ln1))
}

package p1669

import . "leetcode/stucture"

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	i := 0
	left := list1
	for ; i < a-1; i++ {
		left = left.Next
	}
	right := left
	for ; i <= b; i++ {
		right = right.Next
	}
	left.Next = list2
	for list2.Next != nil {
		list2 = list2.Next
	}
	list2.Next = right
	return list1
}

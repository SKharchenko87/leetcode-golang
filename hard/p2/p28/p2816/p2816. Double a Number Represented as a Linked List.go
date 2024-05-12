package p2816

import . "leetcode/stucture"

func doubleIt(head *ListNode) *ListNode {
	if dfs(head) == 1 {
		return &ListNode{Val: 1, Next: head}
	}
	return head
}

func dfs(root *ListNode) int {
	if root == nil {
		return 0
	}
	newVal := root.Val*2 + dfs(root.Next)
	res := newVal / 10
	root.Val = newVal % 10
	return res
}

package p0237

import . "leetcode/stucture"

func run(head *ListNode, node int) *ListNode {
	cur := head
	for cur.Val != node {
		cur = cur.Next
	}
	deleteNode(cur)
	return head
}

func deleteNode(node *ListNode) {
	for node.Next.Next != nil {
		node.Val = node.Next.Val
		node = node.Next
	}
	node.Val = node.Next.Val
	node.Next = nil
}

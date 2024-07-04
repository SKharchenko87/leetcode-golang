package p2181

import . "leetcode/stucture"

func mergeNodes(head *ListNode) *ListNode {
	var dfs func(node *ListNode)
	dfs = func(node *ListNode) {
		if node.Next.Next == nil {
			node.Next = nil
			return
		}
		if node.Next.Val != 0 {
			node.Val += node.Next.Val
			node.Next = node.Next.Next
			dfs(node)
		} else {
			dfs(node.Next)
		}
	}
	dfs(head)
	return head
}

func mergeNodes0(head *ListNode) *ListNode {
	cur := head
	res := head
	for cur != nil && cur.Next != nil {
		if cur.Val == 0 && res.Val != 0 {
			res.Next = cur
			res = cur
		}
		res.Val += cur.Val
		cur = cur.Next
	}
	res.Next = nil
	return head
}

func run(arr []int) []int {
	head := ArrToList(arr)
	mergeNodes(head)
	return ListToArr(head)
}

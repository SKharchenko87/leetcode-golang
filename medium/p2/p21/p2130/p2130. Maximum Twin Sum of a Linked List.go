package p2130

import . "leetcode/stucture"

func pairSum(head *ListNode) int {
	var n int
	cur := head
	for cur != nil {
		cur = cur.Next
		n++
	}
	cur = head
	for i := 0; i < n/2; i++ {
		cur = cur.Next
	}
	var prev, next *ListNode
	for i := 0; i < n/2; i++ {
		next = cur.Next
		cur.Next, cur, prev = prev, next, cur
	}
	var ans int
	for i := 0; i < n/2; i++ {
		ans = max(ans, head.Val+prev.Val)
		head = head.Next
		prev = prev.Next
	}
	return ans
}

func pairSum0(head *ListNode) int {
	index := 0
	arr := make([]int, 100000)
	for head != nil {
		arr[index] = head.Val
		index++
		head = head.Next
	}
	maxSum := 0
	for i := 0; i < index/2; i++ {
		tmp := arr[i] + arr[index-i-1]
		if maxSum < tmp {
			maxSum = tmp
		}
	}
	return maxSum
}

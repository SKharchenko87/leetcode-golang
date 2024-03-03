package p0019

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cur := head
	sz := 0
	for cur != nil {
		cur = cur.Next
		sz++
	}
	i := 1
	var prevNode *ListNode
	cur = head
	for cur != nil {
		if sz-i+1 == n {
			if prevNode == nil {
				head = cur.Next
			} else {
				prevNode.Next = cur.Next
			}
			break
		}
		i++
		prevNode = cur
		cur = cur.Next
	}
	return head
}

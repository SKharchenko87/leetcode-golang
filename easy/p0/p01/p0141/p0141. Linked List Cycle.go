package p0141

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	p1, p2 := head, head
	for p2 != nil && p2.Next != nil {
		p1, p2 = p1.Next, p2.Next.Next
		if p1 == p2 {
			return true
		}
	}
	return false
}

func run(head []int, pos int) bool {
	headList := &ListNode{}
	headList.Val = head[0]
	newList := &ListNode{}
	headList.Next = newList
	var cycle *ListNode
	if pos == 0 {
		cycle = newList
	}
	for i := 1; i < len(head); i++ {
		if pos == i {
			cycle = newList
		}
		newList.Val = head[i]
		newList.Next = &ListNode{}
	}
	newList.Next = cycle
	return hasCycle(headList)
}

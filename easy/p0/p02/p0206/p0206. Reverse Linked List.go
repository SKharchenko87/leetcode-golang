package p0206

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var res *ListNode
	for head != nil {
		res = &ListNode{head.Val, res}
		head = head.Next
	}
	return res
}

func run(headArr []int) []int {
	l := len(headArr)
	if l == 0 {
		return []int{}
	}
	head := new(ListNode)
	save := head
	for i := 0; i < l-1; i++ {
		head.Val = headArr[i]
		next := new(ListNode)
		head.Next = next
		head = next
	}

	head.Val = headArr[l-1]
	head = save
	reverseHead := reverseList(head)
	res := make([]int, len(headArr))
	i := 0
	for reverseHead != nil {
		res[i] = reverseHead.Val
		reverseHead = reverseHead.Next
		i++
	}
	return res
}

package p1171

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	tmp := &ListNode{0, head}
	sumMap := map[int]*ListNode{}
	sumMap[0] = tmp
	sum := 0
	for node := tmp; node != nil; node = node.Next {
		sum += node.Val
		sumMap[sum] = node
	}
	sum = 0
	for node := tmp; node != nil; node = node.Next {
		sum += node.Val
		node.Next = sumMap[sum].Next
	}
	return tmp.Next
}

func arrToList(arr []int) *ListNode {
	l := len(arr)
	if l == 0 {
		return nil
	}
	tmp := &ListNode{}
	res := tmp
	for _, v := range arr {
		tmp.Next = &ListNode{}
		tmp = tmp.Next
		tmp.Val = v
	}
	return res.Next
}

func listToArr(head *ListNode) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func run(head []int) []int {
	h := arrToList(head)
	return listToArr(removeZeroSumSublists(h))
}

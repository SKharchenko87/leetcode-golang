package p0876

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	ans := head
	tmp := head
	for tmp != nil && tmp.Next != nil {
		tmp = tmp.Next.Next
		ans = ans.Next
	}
	return ans
}

func run(arr []int) []int {
	head := &ListNode{}
	tmp := head
	head.Val = arr[0]
	for i := 1; i < len(arr); i++ {
		newList := &ListNode{}
		newList.Val = arr[i]
		head.Next = newList
		head = newList
	}
	head = middleNode(tmp)
	res := make([]int, (len(arr)+1)/2)
	for i := range res {
		fmt.Println(head.Val)
		res[i] = head.Val
		head = head.Next
	}
	return res
}

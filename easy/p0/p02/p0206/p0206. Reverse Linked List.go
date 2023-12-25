package main

import "fmt"

func main() {
	headArr := []int{1, 2, 3, 4, 5}
	head := new(ListNode)
	save := head
	for i, v := range headArr {
		head.Val = v
		fmt.Printf("v: %p %v\n", &head, head)
		if i != len(headArr)-1 {
			next := new(ListNode)
			head.Next = next
			head = next
		}
	}
	var x *ListNode
	x = save
	x = reverseList(x)
	for x != nil {
		fmt.Println(x)
		x = x.Next
	}
}

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

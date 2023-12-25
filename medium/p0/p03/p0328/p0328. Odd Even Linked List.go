package main

import "fmt"

func main() {
	headArr := []int{2, 1, 3, 5, 6, 4, 7}
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
	x = oddEvenList(x)
	for x != nil {
		fmt.Println(x)
		x = x.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	even := &ListNode{head.Val, nil}
	pointEven := even
	head = head.Next
	if head == nil {
		return even
	}
	odd := &ListNode{head.Val, nil}
	pointOdd := odd
	head = head.Next
	index := 0
	for head != nil {
		if index%2 == 0 {
			newNode := &ListNode{head.Val, nil}
			even.Next = newNode
			even = even.Next
		} else {
			newNode := &ListNode{head.Val, nil}
			odd.Next = newNode
			odd = odd.Next
		}
		index++
		head = head.Next
	}
	even.Next = pointOdd
	return pointEven
}

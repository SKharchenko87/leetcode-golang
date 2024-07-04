package stucture

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func SliceToListNode(slice []int) *ListNode {
	lenslice := len(slice)
	p := &(ListNode{Val: slice[lenslice-1], Next: nil})
	for i := lenslice - 2; i >= 0; i-- {
		p = &(ListNode{Val: slice[i], Next: p})
	}
	return p
}

func PrintListNode(ln *ListNode) {
	for ln.Next != nil {
		fmt.Println(ln)
		ln = ln.Next
	}
	fmt.Println(ln)
}

func (ln *ListNode) Print() {
	for {
		fmt.Println(ln.Val)
		if ln.Next == nil {
			break
		}
		ln = ln.Next
	}
}

func (ln *ListNode) ArrayToListNode(arr *[]int) *ListNode {
	return ln.arrayToListNode(arr, 0)
}

func (ln *ListNode) arrayToListNode(arr *[]int, i int) *ListNode {
	if i == len(*arr) {
		return ln
	}
	if i == 0 {
		ln = &ListNode{Val: (*arr)[len(*arr)-1-i], Next: nil}
	} else {
		ln = &ListNode{Val: (*arr)[len(*arr)-1-i], Next: ln}
	}
	i++
	return ln.arrayToListNode(arr, i)
}

func ArrToList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	res := &ListNode{}
	cur := res
	for i, a := range arr {
		cur.Val = a
		if i != len(arr)-1 {
			cur.Next = &ListNode{}
			cur = cur.Next
		}
	}
	return res
}

func ListToArr(head *ListNode) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

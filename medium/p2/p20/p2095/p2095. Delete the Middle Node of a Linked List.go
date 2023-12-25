package main

import "fmt"

func main() {
	//headArr := []int{1, 3, 4, 7, 1, 2, 6}
	//headArr := []int{1, 3, 4, 7}
	//headArr := []int{1, 3}
	headArr := []int{1}
	//headArr := []int{1}
	//headArr := []int{1, 3}
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
	//for x != nil {
	//	fmt.Println(x)
	//	x = x.Next
	//}
	x = deleteMiddle(x)
	for x != nil {
		fmt.Println(x)
		x = x.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (this *ListNode) counterList() int {
	x := this
	count := 0
	for x != nil {
		count++
		x = x.Next
	}
	return count
}

//func counter(this *ListNode) int {
//	x := this
//	count := 0
//	for x != nil {
//		count++
//		x = x.Next
//	}
//	return count
//}
//
//func deleteMiddle(head *ListNode) *ListNode {
//	x := head
//	count := counter(x)
//	if count <= 1 {
//		return nil
//	}
//
//	prevE := head
//	for i := 0; i < count/2-1; i++ {
//		prevE = prevE.Next
//	}
//	prevE.Next = prevE.Next.Next
//
//	return head
//}

//func deleteMiddle2(head *ListNode) *ListNode {
//	Node := head
//	if Node == nil || Node.Next == nil {
//		return nil
//	}
//	var curNode *ListNode
//	for Node != nil {
//		x := Node.Next
//		if x != nil {
//			if curNode == nil {
//				curNode = head
//			} else {
//				curNode = curNode.Next
//			}
//			Node = x.Next
//		} else {
//			Node = x
//		}
//	}
//	curNode.Next = curNode.Next.Next
//	return head
//}

func deleteMiddle(head *ListNode) *ListNode {
	Node := head
	if Node == nil || Node.Next == nil {
		return nil
	}
	Node = Node.Next.Next
	curNode := head
	for Node != nil {
		if Node.Next != nil {
			curNode = curNode.Next
			Node = Node.Next.Next
		} else {
			Node = Node.Next
		}
	}
	curNode.Next = curNode.Next.Next
	return head
}

//  N  x
//  0  1  2  3  4  5  6
//  1, 3, 4, 7, 1, 2, 6
//c

//     N
//  0  1  2  3  4  5  6
//  1, 3, 4, 7, 1, 2, 6
//  c

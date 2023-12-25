package main

import "fmt"

func main() {
	headArr := []int{4, 2, 2, 3}
	//headArr := []int{1, 100000}
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
	print(pairSum(x))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
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

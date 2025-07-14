package p1290

import . "leetcode/stucture"

func getDecimalValue(head *ListNode) (res int) {
	for head != nil {
		res = res<<1 + head.Val
		head = head.Next
	}
	return res
}

func getDecimalValue3(head *ListNode) int {
	var res int
	for head != nil {
		res <<= 1
		res |= head.Val
		head = head.Next
	}
	return res
}

func getDecimalValue2(head *ListNode) int {
	var res int
	countBit := 30
	for head != nil {
		res |= (head.Val << countBit)
		head = head.Next
		countBit--
	}
	res = res >> (countBit + 1)
	return res
}

func getDecimalValue1(head *ListNode) (res int) {
	level := 0
	maxLevel := 0
	var rec func(head *ListNode)
	rec = func(head *ListNode) {
		if head == nil {
			maxLevel = level
			return
		}
		level++
		rec(head.Next)
		res += (head.Val << (maxLevel - level))
		level--
	}
	rec(head)
	return res
}

func getDecimalValue0(head *ListNode) (res int) {
	m := 1
	for head != nil {
		defer func(x int) {
			res += x * m
			m *= 2
		}(head.Val)
		head = head.Next
	}
	return res
}

func run(head []int) int {
	return getDecimalValue(ArrToList(head))
}

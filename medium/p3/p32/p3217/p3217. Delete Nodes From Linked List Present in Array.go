package p3217

import (
	. "leetcode/stucture"
	"slices"
	"sort"
)

func modifiedList(nums []int, head *ListNode) *ListNode {
	template := make([]bool, 100001)
	for i := 0; i < len(nums); i++ {
		template[nums[i]] = true
	}
	head = &ListNode{Val: 0, Next: head}
	cur := head
	for cur.Next != nil {
		if template[cur.Next.Val] {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head.Next
}

func modifiedList3(nums []int, head *ListNode) *ListNode {
	template := make([]bool, 100001)
	for i := 0; i < len(nums); i++ {
		template[nums[i]] = true
	}
	head = &ListNode{Val: 0, Next: head}
	cur := head
	for cur != nil {
		if cur.Next != nil && template[cur.Next.Val] {
			cur.Next = cur.Next.Next
			continue
		}
		cur = cur.Next
	}
	return head.Next
}

func modifiedList2(nums []int, head *ListNode) *ListNode {
	sort.Ints(nums)
	head = &ListNode{Val: 0, Next: head}
	cur := head
	for cur != nil {
		if cur.Next != nil && search(nums, cur.Next.Val) {
			cur.Next = cur.Next.Next
			continue
		}
		cur = cur.Next
	}
	return head.Next
}

func search(nums []int, x int) bool {
	_, res := slices.BinarySearch(nums, x)
	return res
}

func modifiedList1(nums []int, head *ListNode) *ListNode {
	sort.Ints(nums)
	for head != nil && search(nums, head.Val) {
		head = head.Next
	}
	cur := head
	for cur != nil {
		if cur.Next != nil && search(nums, cur.Next.Val) {
			cur.Next = cur.Next.Next
			continue
		}
		cur = cur.Next
	}
	return head
}

func modifiedList0(nums []int, head *ListNode) *ListNode {
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = true
	}
	for head != nil && m[head.Val] {
		head = head.Next
	}
	cur := head
	for cur != nil {
		if cur.Next != nil && m[cur.Next.Val] {
			cur.Next = cur.Next.Next
			continue
		}
		cur = cur.Next
	}
	return head
}

func run(nums []int, root []int) []int {
	return ListToArr(modifiedList(nums, ArrToList(root)))
}

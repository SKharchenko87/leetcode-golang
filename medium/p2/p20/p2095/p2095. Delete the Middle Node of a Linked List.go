package main

import . "leetcode/stucture"

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
			break
		}
	}
	curNode.Next = curNode.Next.Next
	return head
}

func deleteMiddle0(head *ListNode) *ListNode {
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

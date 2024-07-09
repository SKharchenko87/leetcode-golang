package p2058

import (
	. "leetcode/stucture"
	"math"
)

func nodesBetweenCriticalPoints(head *ListNode) []int {
	index := 0
	firstCriticalPointIndex := -1
	for head != nil && head.Next != nil && head.Next.Next != nil {
		isCriticalPoint := head.Val < head.Next.Val && head.Next.Val > head.Next.Next.Val || head.Val > head.Next.Val && head.Next.Val < head.Next.Next.Val
		firstCriticalPointIndex = index
		head = head.Next
		index++
		if isCriticalPoint {
			break
		}
	}
	lastCriticalPointIndex := firstCriticalPointIndex
	minDistance := math.MaxInt
	for head != nil && head.Next != nil && head.Next.Next != nil {
		if head.Val < head.Next.Val && head.Next.Val > head.Next.Next.Val || head.Val > head.Next.Val && head.Next.Val < head.Next.Next.Val {
			minDistance = min(minDistance, index-lastCriticalPointIndex)
			lastCriticalPointIndex = index
		}
		index++
		head = head.Next
	}
	if firstCriticalPointIndex == lastCriticalPointIndex {
		return []int{-1, -1}
	}
	maxDistance := lastCriticalPointIndex - firstCriticalPointIndex
	return []int{minDistance, maxDistance}
}

func nodesBetweenCriticalPoints1(head *ListNode) []int {
	prev1, prev2 := head.Val, head.Val
	index := 0
	firstCriticalPointIndex := -1
	lastCriticalPointIndex := -1
	for head != nil {
		isCriticalPoint := prev2 < prev1 && prev1 > head.Val || prev2 > prev1 && prev1 < head.Val
		firstCriticalPointIndex = index
		lastCriticalPointIndex = index
		prev1, prev2 = head.Val, prev1
		head = head.Next
		index++
		if isCriticalPoint {
			break
		}
	}
	minDistance := math.MaxInt
	for head != nil {
		if prev2 < prev1 && prev1 > head.Val || prev2 > prev1 && prev1 < head.Val {
			minDistance = min(minDistance, index-lastCriticalPointIndex)
			lastCriticalPointIndex = index
		}
		prev1, prev2 = head.Val, prev1
		index++
		head = head.Next
	}
	if firstCriticalPointIndex == lastCriticalPointIndex {
		return []int{-1, -1}
	}
	maxDistance := lastCriticalPointIndex - firstCriticalPointIndex
	return []int{minDistance, maxDistance}
}

func nodesBetweenCriticalPoints0(head *ListNode) []int {
	prev1, prev2 := head.Val, head.Val
	index := 0
	firstCriticalPointIndex := -1
	lastCriticalPointIndex := -1
	minDistance, maxDistance := math.MaxInt, math.MinInt
	for head != nil {
		isCriticalPoint := prev2 < prev1 && prev1 > head.Val || prev2 > prev1 && prev1 < head.Val
		if isCriticalPoint {
			if firstCriticalPointIndex == -1 {
				firstCriticalPointIndex = index
				lastCriticalPointIndex = index
			} else {
				maxDistance = index - firstCriticalPointIndex
				minDistance = min(minDistance, index-lastCriticalPointIndex)
				lastCriticalPointIndex = index
			}
		}
		prev1, prev2 = head.Val, prev1
		index++
		head = head.Next
	}
	if firstCriticalPointIndex == lastCriticalPointIndex {
		return []int{-1, -1}
	}
	return []int{minDistance, maxDistance}
}

func run(arr []int) []int {
	head := ArrToList(arr)
	return nodesBetweenCriticalPoints(head)
}

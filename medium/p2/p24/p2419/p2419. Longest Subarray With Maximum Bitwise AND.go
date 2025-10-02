package p2419

import (
	"container/heap"
	"slices"
)

func longestSubarray(nums []int) int {
	maxElement := 0
	currLen := 0
	result := 0
	for _, num := range nums {
		if num > maxElement {
			maxElement = num
			currLen = 0
			result = 1
		}
		if num == maxElement {
			currLen++
		} else {
			currLen = 0
		}
		result = max(result, currLen)
	}
	return result
}

func longestSubarray7(nums []int) int {
	l := len(nums)
	maxElement := nums[0]
	currLen := 1
	result := 0
	var f = func(x int) {
		if maxElement < x {
			maxElement = x
			result = currLen
		} else if maxElement == x {
			result = max(result, currLen)
		}
	}
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			currLen++
		} else {
			f(nums[i-1])
			currLen = 1
		}
	}
	f(nums[l-1])
	return result
}

func longestSubarray6(nums []int) int {
	l := len(nums)
	maxElement := nums[0]
	currLen := 1
	result := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			currLen++
		} else {
			if maxElement < nums[i-1] {
				maxElement = nums[i-1]
				result = currLen
			} else if maxElement == nums[i-1] {
				result = max(result, currLen)
			}
			currLen = 1
		}
	}
	if maxElement < nums[l-1] {
		maxElement = nums[l-1]
		result = currLen
	} else if maxElement == nums[l-1] {
		result = max(result, currLen)
	}
	return result
}

func longestSubarray5(nums []int) int {
	maxElement := slices.Max(nums)
	currLen := 0
	result := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == maxElement {
			currLen++
		} else {
			result = max(result, currLen)
			currLen = 0
		}
	}
	result = max(result, currLen)
	return result
}

type Element struct {
	num   int
	index int
}

type MaxHeapElement []Element

func (h MaxHeapElement) Len() int {
	return len(h)
}
func (h MaxHeapElement) Less(i, j int) bool {
	if h[i].num == h[j].num {
		return h[i].index < h[j].index
	}
	return h[i].num > h[j].num
}

func (h MaxHeapElement) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MaxHeapElement) Push(x interface{}) {
	*h = append(*h, x.(Element))
}

func (h *MaxHeapElement) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func longestSubarray4(nums []int) int {
	l := len(nums)
	mhe := make(MaxHeapElement, l)
	heap.Init(&mhe)
	for i := 0; i < l; i++ {
		heap.Push(&mhe, Element{nums[i], i})
	}
	prev := 0
	currElement := heap.Pop(&mhe).(Element)
	curr := currElement.num
	indexPrev := currElement.index
	currLen := 1
	result := currLen
	for mhe.Len() > 0 {
		prev = curr
		currElement = heap.Pop(&mhe).(Element)
		curr = currElement.num
		if prev > curr {
			break
		}
		index := currElement.index
		if index-1 == indexPrev {
			currLen++
		} else {
			result = max(result, currLen)
			currLen = 1
		}
		indexPrev = index
	}
	result = max(result, currLen)
	return result
}

func longestSubarray3(nums []int) int {
	result := 0
	maxNum := 0
	for i := 0; i < len(nums); i++ {
		currLen := 1
		currNum := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if currNum <= currNum&nums[j] {
				currLen++
				currNum &= nums[j]
			} else {
				break
			}
		}
		if maxNum <= currNum {
			if maxNum == currNum {
				result = max(result, currLen)
			} else {
				result = currLen
			}
			maxNum = currNum
		}
	}
	return result
}

/*
=
=
=
=
=
=
=
=
=
=
=
=
=
=
=
=
=
=
=
=
=
=
=
=
=
=
=
=
=
=
*/

func longestSubarray2(nums []int) int {
	maxV := nums[0]
	l := len(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxV {
			maxV = nums[i]
		}
	}
	maxLength := 0
	for i := 0; i < l; i++ {
		if nums[i] == maxV {
			j := i + 1
			for ; j < l && nums[i] == nums[j]; j++ {
			}
			if maxLength < j-i {
				maxLength = j - i
			}
			i = j - 1
		}
	}
	return maxLength
}

func longestSubarray1(nums []int) int {
	maxV := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxV {
			maxV = nums[i]
		}
	}
	currentLength := 0
	maxLength := 0
	for i := range nums {
		if nums[i] == maxV {
			currentLength++
		} else {
			currentLength = 0
		}
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}
	return maxLength
}

func longestSubarray0(nums []int) int {
	currentV := nums[0]
	maxV := nums[0]
	maxLastIndex := 0
	currentLength := 1
	maxLength := 1
	for i := 1; i < len(nums); i++ {
		currentV = nums[i]
		if currentV > maxV {
			maxV = currentV
			currentLength = 1
			maxLength = 1
			maxLastIndex = i
		} else if currentV == maxV {
			if maxLastIndex+1 == i {
				currentLength++
			} else {
				currentLength = 1
			}
			maxLastIndex = i
		}
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}
	return maxLength
}

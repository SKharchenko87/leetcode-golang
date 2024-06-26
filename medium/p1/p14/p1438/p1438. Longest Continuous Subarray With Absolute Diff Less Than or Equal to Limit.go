package p1438

func longestSubarray(nums []int, limit int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	l := len(nums)
	maxQ, maxQLeft, maxQRight := make([]int, l), 0, -1
	minQ, minQLeft, minQRight := make([]int, l), 0, -1

	i := 0

	for _, element := range nums {
		for maxQLeft <= maxQRight && element > maxQ[maxQRight] {
			maxQRight--
		}
		for minQLeft <= minQRight && element < minQ[minQRight] {
			minQRight--
		}
		minQRight++
		maxQRight++
		maxQ[maxQRight] = element
		minQ[minQRight] = element

		if maxQ[maxQLeft]-minQ[minQLeft] > limit {
			if maxQ[maxQLeft] == nums[i] {
				maxQLeft++
			}
			if minQ[minQLeft] == nums[i] {
				minQLeft++
			}
			i++
		}
	}

	return len(nums) - i
}

func longestSubarray0(nums []int, limit int) int {
	l := len(nums)
	res := 0
	for i := 0; i < l; i++ {
		currentLength := 0
		currentMax, currentMaxIndex := nums[i], i
		currentMin, currentMinIndex := nums[i], i
		for j := i; j < l; j++ {
			if abs(currentMax-nums[j]) > limit || abs(currentMin-nums[j]) > limit {
				if abs(currentMax-nums[j]) > limit && abs(currentMin-nums[j]) > limit {
					i = max(currentMaxIndex, currentMinIndex)
				} else if abs(currentMax-nums[j]) > limit {
					i = currentMaxIndex
				} else {
					i = currentMinIndex
				}
				break
			}
			if currentMax <= nums[j] {
				currentMax, currentMaxIndex = nums[j], j
			}
			if currentMin >= nums[j] {
				currentMin, currentMinIndex = nums[j], j
			}
			currentLength++
			if currentLength > res {
				res = currentLength
			}
		}
		if res+i >= l {
			break
		}
	}
	return res
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -1 * x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

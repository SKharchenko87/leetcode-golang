package p1695

func maximumUniqueSubarray(nums []int) int {
	numLastIndex := [10001]int{}
	start, end := 0, 0
	currentSum := 0
	maxSum := currentSum
	for ; end < len(nums); end++ {
		num := nums[end]

		for lastIndex := numLastIndex[num] - 1; lastIndex >= start; start++ {
			currentSum -= nums[start]
		}

		numLastIndex[num] = end + 1
		currentSum += num
		maxSum = max(maxSum, currentSum)
	}
	return maxSum
}

func maximumUniqueSubarray3(nums []int) int {
	numLastIndex := [10001]int{}
	for i := 0; i < len(numLastIndex); i++ {
		numLastIndex[i] = -1
	}
	start, end := 0, 0
	currentSum := 0
	maxSum := currentSum
	for ; end < len(nums); end++ {
		num := nums[end]

		for lastIndex := numLastIndex[num]; lastIndex >= start; start++ {
			currentSum -= nums[start]
		}

		numLastIndex[num] = end
		currentSum += num
		maxSum = max(maxSum, currentSum)
	}
	return maxSum
}

func maximumUniqueSubarray2(nums []int) int {
	numLastIndex := map[int]int{}
	start, end := 0, 0
	currentSum := 0
	maxSum := 0
	for ; end < len(nums); end++ {
		num := nums[end]
		if index, ok := numLastIndex[num]; ok && index >= start {
			for ; index >= start; start++ {
				currentSum -= nums[start]
			}
		}
		numLastIndex[num] = end
		currentSum += num
		maxSum = max(maxSum, currentSum)
	}
	return maxSum
}

func maximumUniqueSubarray1(nums []int) int {
	prevSum := make([]int, len(nums)+1)
	prevSum[0] = 0

	m := make(map[int]int)
	res := 0
	lastIndex := 0

	for i := 0; i < len(nums); i++ {
		prevSum[i+1] = prevSum[i] + nums[i]
		res = max(res, prevSum[i]-prevSum[lastIndex])
		if index, exists := m[nums[i]]; exists {
			lastIndex = max(lastIndex, index+1)
		}
		m[nums[i]] = i
	}

	res = max(res, prevSum[len(nums)]-prevSum[lastIndex])
	return res
}

func maximumUniqueSubarray0(nums []int) int {
	prevSum := make([]int, len(nums)+1)
	prevSum[0] = 0
	for i := 0; i < len(nums); i++ {
		prevSum[i+1] = prevSum[i] + nums[i]
	}
	m := make(map[int]int)
	res := 0
	lastIndex := 0
	for i := 0; i < len(nums); i++ {
		res = max(res, prevSum[i]-prevSum[lastIndex])
		if index, exists := m[nums[i]]; exists {
			lastIndex = max(lastIndex, index+1)
		}
		m[nums[i]] = i
	}
	res = max(res, prevSum[len(nums)]-prevSum[lastIndex])
	return res
}

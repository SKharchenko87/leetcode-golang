package p0396

func maxRotateFunction(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}
	totalSum := 0
	currentF := 0
	for i := 0; i < n; i++ {
		totalSum += nums[i]
		currentF += i * nums[i]
	}

	maxF := currentF
	for i := 1; i < n; i++ {
		currentF = currentF + totalSum - n*nums[n-i]
		maxF = max(maxF, currentF)

	}

	return maxF
}

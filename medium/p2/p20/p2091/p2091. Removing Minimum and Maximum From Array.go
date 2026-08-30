package p2091

func minimumDeletions(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 1
	} else if n <= 3 {
		return 2
	}
	minVal, maxVal := nums[0], nums[0]
	minValIndex, maxValIndex := 0, 0
	for i := 1; i < n; i++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
			maxValIndex = i
		}
		if nums[i] < minVal {
			minVal = nums[i]
			minValIndex = i
		}
	}
	l, r := min(maxValIndex, minValIndex), max(maxValIndex, minValIndex)
	return min(r, n-l-1, l+n-r) + 1
}

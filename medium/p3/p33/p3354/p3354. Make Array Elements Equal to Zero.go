package p3354

func countValidSelections(nums []int) int {
	l := len(nums)
	right, left := make([]int, l), make([]int, l)
	left[l-1] = nums[l-1]
	for i := l - 2; i >= 0; i-- {
		left[i] = nums[i] + left[i+1]
	}
	right[0] = nums[0]
	for i := 1; i < l; i++ {
		right[i] = right[i-1] + nums[i]
	}
	res := 0
	for i := 0; i < l-1; i++ {
		if nums[i] == 0 {
			if right[i] == left[i+1] {
				res += 2
			} else if abs(right[i]-left[i+1]) == 1 {
				res += 1
			}
		}
	}
	if nums[0] == 0 && left[0] == 0 {
		res++
	}
	if nums[l-1] == 0 && right[l-1] == 0 || nums[l-1] == 0 && right[l-1] == 1 {
		res++
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

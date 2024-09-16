package p3

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

func abs64(x int64) int64 {
	if x >= 0 {
		return x
	}
	return -1 * x

}

func min64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func max64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func maxValue(nums []int, k int) int {
	l := len(nums)
	bitsL, bitsR := make([]int, 7), make([]int, 7)

	left, right := make([]int, l), make([]int, l)
	left[0] = nums[0]
	for i := 1; i < l; i++ {
		left[i] = left[i-1] | nums[i]
	}
	right[l-1] = nums[l-1]
	for i := l - 2; i >= 0; i-- {
		right[i] = right[i+1] | nums[i]
	}
	max

}

// 2= 010
// 4= 100
// 5= 101
// 6= 110
// 7= 111

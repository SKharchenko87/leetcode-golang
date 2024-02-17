package p1

func maxOperations(nums []int) int {
	cnt := 1
	nums[1] += nums[0]
	for i := 2; i < len(nums)-1; i += 2 {
		nums[i+1] += nums[i]
		if nums[i+1] != nums[i-1] {
			break
		}
		cnt++
	}
	return cnt
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

package p3105

func longestMonotonicSubarray(nums []int) int {
	l := len(nums)
	if l < 2 {
		return l
	}

	ascLen := 1
	descLen := 1
	res := 0
	for i := 1; i < l; i++ {
		diff := nums[i] - nums[i-1]
		if diff == 0 {
			ascLen = 1
			descLen = 1
		} else if diff < 0 {
			ascLen = 1
			descLen++
		} else {
			ascLen++
			descLen = 1
		}
		res = max(res, ascLen, descLen)
	}

	return res
}

func longestMonotonicSubarray0(nums []int) int {
	l := len(nums)
	if l < 2 {
		return l
	}

	maxLength := 1
	curLength := 1
	prev := nums[1] - nums[0]
	for i := 1; i < l; i++ {
		tmp := nums[i] - nums[i-1]
		if prev > 0 && tmp > 0 || prev < 0 && tmp < 0 {
			curLength++
		} else {
			if curLength > maxLength {
				maxLength = curLength
			}
			if tmp == 0 {
				curLength = 1
			} else {
				curLength = 2
			}
		}
		prev = tmp
	}

	if curLength > maxLength {
		maxLength = curLength
	}

	return maxLength
}

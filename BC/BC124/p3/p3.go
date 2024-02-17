package p3

//
//1:{ 5:[2,5], 7:[0,3], 7:[1,4]}
//2:{ 5:[3,4]}
//3:{ 5:[3,4]}

func maxOperations(nums []int) int {
	count := 0
	var llarr, rrarr, lrarr [3]int

	m := make([][][3]int, 2)
	m[0] = append(m[0], ll(nums, 0, len(nums)-1))
	m[0] = append(m[0], rr(nums, 0, len(nums)-1))
	m[0] = append(m[0], lr(nums, 0, len(nums)-1))
	count++
	for count <= len(nums)/2 {
		arr := m[0]
		for _, tripl := range arr {
			if tripl[0] != -1 {
				llarr = ll(nums, tripl[1], tripl[2])
				rrarr = rr(nums, tripl[1], tripl[2])
				lrarr = lr(nums, tripl[1], tripl[2])
				if tripl[0] == llarr[0] {
					m[1] = append(m[1], llarr)
				}
				if tripl[0] == rrarr[0] {
					m[1] = append(m[1], rrarr)
				}
				if tripl[0] == lrarr[0] {
					m[1] = append(m[1], lrarr)
				}
			}
		}
		if len(m[1]) == 0 {
			return count
		}
		count++
		m = append(m[1:], [][3]int{})
	}
	return len(m)
}

func ll(nums []int, start, end int) [3]int {
	if start < end {
		return [3]int{nums[start] + nums[start+1], start + 2, end}
	} else {
		return [3]int{-1, 0, 0}
	}
}

func rr(nums []int, start, end int) [3]int {
	if start < end {
		return [3]int{nums[end-1] + nums[end], start, end - 2}
	} else {
		return [3]int{-1, 0, 0}
	}
}

func lr(nums []int, start, end int) [3]int {
	if start < end {
		return [3]int{nums[start] + nums[end], start + 1, end - 1}
	} else {
		return [3]int{-1, 0, 0}
	}
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

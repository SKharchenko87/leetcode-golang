package p0494

func findTargetSumWays(nums []int, target int) int {
	l := byte(len(nums))
	if l == 0 {
		return 0
	}
	m := map[int]int{0: 1}
	nm := map[int]int{}

	for level := byte(0); level < l; level++ {
		for i, c := range m {
			nm[i+nums[level]] += c
			nm[i-nums[level]] += c
		}
		//clear(m)
		m = map[int]int{}
		nm, m = m, nm
	}

	return m[target]
}

func findTargetSumWays1(nums []int, target int) int {
	l := byte(len(nums))
	if l == 0 {
		return 0
	}
	stackVal := make([]int, 0, 1<<l)
	stackVal = append(stackVal, nums[0])
	stackVal = append(stackVal, -nums[0])
	stackLevel := make([]byte, 0, 1<<l)
	stackLevel = append(stackLevel, 0, 0)
	res := 0
	i := 0
	for ; i < len(stackVal); i++ {
		curVal := stackVal[i]
		curLevel := stackLevel[i]
		nextLevel := curLevel + 1
		if curLevel != l-1 {
			stackVal = append(stackVal, curVal+nums[nextLevel])
			stackLevel = append(stackLevel, nextLevel)
			stackVal = append(stackVal, curVal-nums[nextLevel])
			stackLevel = append(stackLevel, nextLevel)
		} else if curVal == target {
			res++
		}
	}
	return res
}

// TLE
func findTargetSumWays0(nums []int, target int) int {
	l := byte(len(nums))
	if l == 0 {
		return 0
	}
	stackVal := []int{nums[0], -nums[0]}
	stackLevel := []byte{0, 0}
	res := 0
	for len(stackVal) > 0 {
		curVal := stackVal[0]
		stackVal = stackVal[1:]
		curLevel := stackLevel[0]
		stackLevel = stackLevel[1:]
		nextLevel := curLevel + 1
		if curLevel != l-1 {
			stackVal = append(stackVal, curVal+nums[nextLevel])
			stackLevel = append(stackLevel, nextLevel)
			stackVal = append(stackVal, curVal-nums[nextLevel])
			stackLevel = append(stackLevel, nextLevel)
		} else if curVal == target {
			res++
		}
	}
	return res
}

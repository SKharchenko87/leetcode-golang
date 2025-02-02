package p1752

func check(nums []int) bool {
	l := len(nums)
	cntSplit := 0
	if nums[0] >= nums[l-1] {
		cntSplit = 1
	}
	for i := 0; i < l-1; i++ {
		if nums[i] > nums[i+1] {
			cntSplit--
			if cntSplit < 0 {
				return false
			}
		}
	}
	return true
}

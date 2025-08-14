package p2134

func minSwaps(nums []int) int {
	l := len(nums)
	cnt := 0
	for i := 0; i < l; i++ {
		if nums[i] == 1 {
			cnt++
		}
	}
	if cnt == 0 || cnt == l {
		return 0
	}
	cnt0 := 0
	for i := 0; i < cnt; i++ {
		if nums[i] == 0 {
			cnt0++
		}
	}
	res := cnt0
	for i := cnt; i < l+cnt; i++ {
		if nums[(i-cnt)%l] == 0 {
			cnt0--
		}
		if nums[i%l] == 0 {
			cnt0++
		}
		if res > cnt0 {
			res = cnt0
		}
	}
	return res
}

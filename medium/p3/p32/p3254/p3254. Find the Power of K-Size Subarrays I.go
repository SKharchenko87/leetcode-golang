package p3254

func resultsArray(nums []int, k int) []int {
	l := len(nums)
	if k == 1 {
		return nums
	}
	res := make([]int, l-k+1)
	cnt := 0
	for i := 1; i < k; i++ {
		if nums[i]-nums[i-1] != 1 {
			cnt = k + 1
		}
		if cnt > 0 {
			cnt--
		}
	}

	for i := k - 1; i < l; i++ {
		if nums[i]-nums[i-1] != 1 {
			cnt = k
		}
		if cnt > 0 {
			cnt--
		}
		if cnt == 0 {
			res[i-k+1] = nums[i]
		} else {
			res[i-k+1] = -1
		}
	}
	return res
}

package p2537

func countGood(nums []int, k int) int64 {
	l := len(nums)
	m := make(map[int]int)
	m[nums[0]]++
	var res int64
	var windowSum int
	var i, j int
	j = 1
	for ; i < l-1; i++ {
		for ; j < l && windowSum < k; j++ {
			windowSum = windowSum + m[nums[j]]
			if windowSum >= k {
				res += int64(l - j)
			}
			m[nums[j]]++
		}
		m[nums[i]]--
		windowSum = windowSum - m[nums[i]]
		if windowSum >= k {
			res += int64(l - j + 1)
		}
	}
	return res
}

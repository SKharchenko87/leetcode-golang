package p2364

func countBadPairs(nums []int) int64 {
	l := len(nums)
	var countPairs int64 = int64(l * (l - 1) / 2)
	m := make(map[int]int)
	for i := 0; i < l; i++ {
		m[nums[i]-i]++
		countPairs -= int64(m[nums[i]-i] - 1)
	}
	return countPairs
}

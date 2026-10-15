package p2913

func sumCounts(nums []int) int {
	res := 0
	n := len(nums)
	m := make(map[int]int, n)
	for i := 0; i < len(nums); i++ {
		clear(m)
		for j := i; j < len(nums); j++ {
			m[nums[j]]++
			res += len(m) * len(m)
		}
	}
	return res
}

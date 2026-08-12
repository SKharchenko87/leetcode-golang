package p2958

func maxSubarrayLength(nums []int, k int) int {
	i, j, n := 0, 0, len(nums)
	m := make(map[int]int, n)
	res := 0
	for ; i < n; i++ {
		vi := nums[i]
		m[vi]++
		for ; m[vi] > k; j++ {
			m[nums[j]]--
		}
		res = max(res, i-j+1)
	}
	return res
}

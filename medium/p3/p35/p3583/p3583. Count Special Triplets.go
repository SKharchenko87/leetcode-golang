package p3583

const MOD = 1e9 + 7

var left, right = make([]int, 100001), make([]int, 100001)

func specialTriplets(nums []int) int {
	l := len(nums)
	clear(left)
	for i := 2; i < l; i++ {
		right[nums[i]]++
	}
	left[nums[0]]++
	res := 0
	for i := 1; i < l-1; i++ {
		mid := nums[i] * 2
		if mid <= 100001 {
			lr := left[mid] * right[mid]
			res = (res + lr) % MOD
		}
		left[nums[i]]++
		right[nums[i+1]]--
	}
	return res
}

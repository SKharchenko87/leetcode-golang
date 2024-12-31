package p0689

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	i := 0
	w := 0
	for ; i < k; i++ {
		w += nums[i]
		nums[i] = w
	}
	for ; i < len(nums); i++ {
		w += nums[i] - nums[i-k]
		nums[i] = w
	}
	return nums
}

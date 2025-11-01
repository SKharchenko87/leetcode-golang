package p3350

func maxIncreasingSubarrays(nums []int) int {
	l := len(nums)
	prev := 0
	cur := 1
	ans := 1
	for i := 1; i < l; i++ {
		if nums[i] > nums[i-1] {
			cur++
		} else {
			prev = cur
			cur = 1
		}
		ans = max(ans, cur/2)
		ans = max(ans, min(cur, prev))
	}
	return ans
}

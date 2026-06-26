package p3739

/*ToDo*/
func countMajoritySubarrays(nums []int, target int) int64 {
	n := len(nums)
	// represents the occurrence count of prefix sums -n, -(n-1), ..., 0, 1,
	// ..., n, with index offset by n.
	pre := make([]int, n*2+1)
	pre[n] = 1
	cnt := n
	var ans, presum int64 = 0, 0
	for i := 0; i < n; i++ {
		if nums[i] == target {
			presum += int64(pre[cnt])
			cnt++
			pre[cnt]++
		} else {
			cnt--
			presum -= int64(pre[cnt])
			pre[cnt]++
		}
		ans += presum
	}
	return ans
}

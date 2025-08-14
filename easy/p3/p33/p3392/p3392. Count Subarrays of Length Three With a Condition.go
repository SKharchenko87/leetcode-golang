package p3392

func countSubarrays(nums []int) int {
	res := 0
	for i := 1; i < len(nums)-1; i++ {
		if nums[i]%2 == 0 && nums[i]/2 == nums[i-1]+nums[i+1] {
			res++
		}
	}
	return res
}

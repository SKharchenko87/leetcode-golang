package p2560

func minCapability(nums []int, k int) int {
	check := func(x int) bool {
		take := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] <= x {
				take++
				i++
			}
		}
		return take >= k
	}
	left := 0
	right := 1000_000_000
	for left < right {
		mid := left + (right-left)/2
		if !check(mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

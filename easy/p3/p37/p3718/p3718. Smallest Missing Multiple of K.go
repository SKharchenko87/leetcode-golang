package p3718

func missingMultiple(nums []int, k int) int {
	counter := [101]bool{}
	for i := 0; i < len(nums); i++ {
		if nums[i]%k == 0 {
			counter[nums[i]/k] = true
		}
	}
	for i := 1; i < 101; i++ {
		if !counter[i] {
			return k * i
		}
	}
	return 101
}

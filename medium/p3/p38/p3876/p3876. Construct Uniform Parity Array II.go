package p3876

func uniformArray(nums1 []int) bool {
	oddCount := 0
	evenCount := 0
	minVal := 1<<32 - 1
	n := len(nums1)
	for i := 0; i < n; i++ {
		if nums1[i]%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
		minVal = min(minVal, nums1[i])
	}
	return minVal%2 == 1 || evenCount == n || oddCount == n || n > 2 && !(minVal%2 == 0 && evenCount > 0 && oddCount > 0)
}

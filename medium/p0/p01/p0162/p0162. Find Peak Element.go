package p0162

// ToDo решение позаимствона, нужно разобраться как это работает
func findPeakElement(nums []int) int {
	start := 0
	end := len(nums) - 1
	for start < end {
		mid := start + (end-start)/2
		if nums[mid] < nums[mid+1] {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return end
}

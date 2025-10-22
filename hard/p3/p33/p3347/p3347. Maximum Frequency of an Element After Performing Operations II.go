package p3347

import "sort"

/*ToDo*/
func maxFrequency(nums []int, k int, numOperations int) int {
	l := len(nums)
	sort.Ints(nums)
	left := 0
	right := 0
	result := 1
	var curFreq, curIndex int
	for ; left < l; left++ {
		curFreq = 0
		for curIndex = left; curIndex < l && nums[left] == nums[curIndex]; curIndex++ {
			curFreq++
		}
		for ; right < l && nums[left]+2*k > nums[right]; right++ {
		}
		result = max(result, min(right-left+1, numOperations+curFreq))
	}
	return result
}

/*TLE*/
func maxFrequency0(nums []int, k int, numOperations int) int {
	l := len(nums)
	sort.Ints(nums)
	start := nums[0] - k
	end := nums[l-1] + k + 1
	left := 0
	right := 0
	result := -1
	cur := 0
	curFreq := 0
	for i := start; i < end; i++ {
		curFreq = 0
		for ; cur < l && nums[cur] <= i; cur++ {
			if nums[cur] == i {
				curFreq++
			}
		}
		for ; right < l && nums[right] <= i+k; right++ {
		}
		for ; left < right && nums[left] < i-k; left++ {
		}
		result = max(result, min(right-left-curFreq, numOperations)+curFreq)
	}
	return result
}

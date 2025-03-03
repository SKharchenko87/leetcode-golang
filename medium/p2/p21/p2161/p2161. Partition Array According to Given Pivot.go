package p2161

func pivotArray(nums []int, pivot int) []int {
	l := len(nums)
	res := make([]int, l)
	leftIndex := 0
	for i := 0; i < l; i++ {
		if nums[i] < pivot {
			res[leftIndex] = nums[i]
			leftIndex++
		}
	}
	rightIndex := l - 1
	for i := l - 1; i >= 0 && leftIndex <= rightIndex; i-- {
		if nums[i] > pivot {
			res[rightIndex] = nums[i]
			rightIndex--
		}
	}
	for i := leftIndex; i <= rightIndex; i++ {
		res[i] = pivot
	}
	return res
}

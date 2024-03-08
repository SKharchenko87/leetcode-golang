package p1

func resultArray(nums []int) []int {
	l := len(nums)
	arr1, arr2 := make([]int, 0, (l+1)/2), make([]int, 0, (l+1)/2)
	i := 2
	arr1, arr2 = append(arr1, nums[0]), append(arr2, nums[1])
	index1, index2 := 0, 0
	for ; i < l; i++ {
		if arr1[index1] > arr2[index2] {
			arr1 = append(arr1, nums[i])
			index1++
		} else {
			arr2 = append(arr2, nums[i])
			index2++
		}
	}
	return append(arr1, arr2...)
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -1 * x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs64(x int64) int64 {
	if x >= 0 {
		return x
	}
	return -1 * x

}

func min64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func max64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

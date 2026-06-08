package p2161

func pivotArray(nums []int, pivot int) []int {
	l := len(nums)
	res := make([]int, 0, l)
	cntPivot := 0
	for i := 0; i < l; i++ {
		if nums[i] < pivot {
			res = append(res, nums[i])
		} else if nums[i] == pivot {
			cntPivot++
		}
	}
	for i := 0; i < cntPivot; i++ {
		res = append(res, pivot)
	}
	for i := 0; i < l; i++ {
		if nums[i] > pivot {
			res = append(res, nums[i])
		}
	}
	return res
}

func pivotArray1(nums []int, pivot int) []int {
	l := len(nums)
	less := make([]int, 0, l/3)
	equal := make([]int, 0, l/3)
	greater := make([]int, 0, l/3)
	for i := 0; i < l; i++ {
		if nums[i] < pivot {
			less = append(less, nums[i])
		} else if nums[i] > pivot {
			greater = append(greater, nums[i])
		} else {
			equal = append(equal, nums[i])
		}
	}
	less = append(less, equal...)
	less = append(less, greater...)
	return less
}

func pivotArray0(nums []int, pivot int) []int {
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

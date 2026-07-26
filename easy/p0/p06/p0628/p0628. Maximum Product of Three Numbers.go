package p0628

func maximumProduct(nums []int) int {
	min0, min1 := nums[0], 10001
	max0, max1, max2 := nums[0], -10001, -10001
	for i := 1; i < len(nums); i++ {
		switch {
		case min0 >= nums[i]:
			min0, min1 = nums[i], min0
		case min1 >= nums[i]:
			min1 = nums[i]
		}
		switch {
		case max0 <= nums[i]:
			max0, max1, max2 = nums[i], max0, max1
		case max1 <= nums[i]:
			max1, max2 = nums[i], max1
		case max2 <= nums[i]:
			max2 = nums[i]
		}
	}
	return max(max0*min0*min1, max0*max1*max2)
}

func maximumProduct0(nums []int) int {
	min0, min1, min2 := 10001, 10001, 10001
	max0, max1, max2 := -10001, -10001, -10001
	for i := 0; i < len(nums); i++ {
		if min0 >= nums[i] {
			min0, min1, min2 = nums[i], min0, min1
		} else if min1 >= nums[i] {
			min1, min2 = nums[i], min1
		} else if min2 >= nums[i] {
			min2 = nums[i]
		}

		if max0 <= nums[i] {
			max0, max1, max2 = nums[i], max0, max1
		} else if max1 <= nums[i] {
			max1, max2 = nums[i], max1
		} else if max2 <= nums[i] {
			max2 = nums[i]
		}
	}

	return max(max0*max1*max2, min0*min1*max0)
}

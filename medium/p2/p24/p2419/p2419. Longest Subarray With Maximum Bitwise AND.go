package p2419

func longestSubarray(nums []int) int {
	maxV := nums[0]
	l := len(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxV {
			maxV = nums[i]
		}
	}
	maxLength := 0
	for i := 0; i < l; i++ {
		if nums[i] == maxV {
			j := i + 1
			for ; j < l && nums[i] == nums[j]; j++ {
			}
			if maxLength < j-i {
				maxLength = j - i
			}
			i = j - 1
		}
	}
	return maxLength
}

func longestSubarray1(nums []int) int {
	maxV := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxV {
			maxV = nums[i]
		}
	}
	currentLength := 0
	maxLength := 0
	for i := range nums {
		if nums[i] == maxV {
			currentLength++
		} else {
			currentLength = 0
		}
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}
	return maxLength
}

func longestSubarray0(nums []int) int {
	currentV := nums[0]
	maxV := nums[0]
	maxLastIndex := 0
	currentLength := 1
	maxLength := 1
	for i := 1; i < len(nums); i++ {
		currentV = nums[i]
		if currentV > maxV {
			maxV = currentV
			currentLength = 1
			maxLength = 1
			maxLastIndex = i
		} else if currentV == maxV {
			if maxLastIndex+1 == i {
				currentLength++
			} else {
				currentLength = 1
			}
			maxLastIndex = i
		}
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}
	return maxLength
}

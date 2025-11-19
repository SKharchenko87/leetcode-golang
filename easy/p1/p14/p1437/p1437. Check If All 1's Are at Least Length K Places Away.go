package p1437

func kLengthApart(nums []int, k int) bool {
	k++
	prevOneIndex := -k
	for i, num := range nums {
		if num == 1 {
			if i < k+prevOneIndex {
				return false
			}
			prevOneIndex = i
		}
	}
	return true
}

func kLengthApart2(nums []int, k int) bool {
	prevOneIndex := -k - 1
	for i, num := range nums {
		if num == 1 {
			if i-prevOneIndex < k+1 {
				return false
			}
			prevOneIndex = i
		}
	}
	return true
}

func kLengthApart1(nums []int, k int) bool {
	prevOneIndex := -k - 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			if i-prevOneIndex < k+1 {
				return false
			}
			prevOneIndex = i
		}
	}
	return true
}

func kLengthApart0(nums []int, k int) bool {
	cur0 := k
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			cur0++
		} else {
			if cur0 < k {
				return false
			}
			cur0 = 0
		}
	}
	return true
}

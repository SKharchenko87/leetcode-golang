package p1760

func minimumSize(nums []int, maxOperations int) int {
	//right, left := nums[0], nums[0]
	right := nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		if right < nums[i] {
			right = nums[i]
		}
		//if left > nums[i] {
		//	left = nums[i]
		//}
	}
	left := sum / (len(nums) + maxOperations)
	if left == 0 {
		left++
	}
	mid := right
	for left < right {
		mid = (left + right) / 2
		if check(nums, mid, maxOperations) {
			right = mid
		} else {
			left = mid + 1
		}

	}
	return left
}

func check(nums []int, maxNum int, maxOperations int) bool {
	for _, num := range nums {
		if num > maxNum {
			maxOperations -= num/maxNum - 1
			if num%maxNum > 0 {
				maxOperations--
			}
			if maxOperations < 0 {
				return false
			}
		}

	}
	return true
}

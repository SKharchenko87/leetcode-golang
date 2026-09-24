package p3550

func smallestIndex(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if i == digitSum(nums[i]) {
			return i
		}
	}
	return -1
}

func digitSum(d int) (res int) {
	for d > 0 {
		res += d % 10
		d /= 10
	}
	return
}

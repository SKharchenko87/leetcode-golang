package p2918

func minSum(nums1 []int, nums2 []int) int64 {
	sum1, cntZero1 := calc(nums1)
	sum2, cntZero2 := calc(nums2)
	sum := max(sum1+int64(cntZero1), sum2+int64(cntZero2))
	if (cntZero1 == 0 && sum1 < sum) || (cntZero2 == 0 && sum2 < sum) {
		return -1
	}
	return sum
}

func calc(nums []int) (sum int64, cnt int) {
	for i := 0; i < len(nums); i++ {
		sum += int64(nums[i])
		if nums[i] == 0 {
			cnt += 1
		}
	}
	return
}

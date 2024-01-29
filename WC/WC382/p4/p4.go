package p4

func minOrAfterOperations(nums []int, k int) int {
	n := len(nums)
	arr := [30]int{}
	for _, v := range nums {
		for i := 0; v > 0; i++ {
			arr[i] += v % 2
			v = v / 2
		}
	}

	return 0
}

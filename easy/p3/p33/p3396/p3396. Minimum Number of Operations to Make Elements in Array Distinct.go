package p3396

func minimumOperations(nums []int) int {
	freq := make([]int, 101)
	l := len(nums)
	for i := l - 1; i >= 0; i-- {
		n := nums[i]
		freq[n]++
		if freq[n] == 2 {
			return i/3 + 1
		}
	}
	return 0
}

func minimumOperations1(nums []int) int {
	freq := make(map[int]int, 101)
	l := len(nums)
	for i := l - 1; i >= 0; i-- {
		n := nums[i]
		freq[n]++
		if freq[n] == 2 {
			return i/3 + 1
		}
	}
	return 0
}

func minimumOperations0(nums []int) int {
	freq := make(map[int]int, 101)
	cntNotDistinct := 0
	for _, n := range nums {
		freq[n]++
		if freq[n] == 2 {
			cntNotDistinct++
		}
	}
	if cntNotDistinct == 0 {
		return 0
	}
	for i, n := range nums {
		if freq[n] == 2 {
			cntNotDistinct--
			if cntNotDistinct == 0 {
				return i/3 + 1
			}
		}
		freq[n]--
	}
	return 0
}

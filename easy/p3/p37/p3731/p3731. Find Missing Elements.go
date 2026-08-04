package p3731

func findMissingElements(nums []int) []int {
	everyInteger := [2]int64{}
	minVal, maxVal := 101, 0
	for _, v := range nums {
		everyInteger[v/64] |= 1 << (v % 64)
		minVal = min(minVal, v)
		maxVal = max(maxVal, v)
	}
	res := make([]int, 0, maxVal-minVal)
	for i := minVal + 1; i < maxVal; i++ {
		if everyInteger[i/64]&(1<<(i%64)) == 0 {
			res = append(res, i)
		}
	}
	return res
}

func findMissingElements0(nums []int) []int {
	everyInteger := make([]bool, 101)
	minVal, maxVal := 101, 0
	for _, v := range nums {
		everyInteger[v] = true
		minVal = min(minVal, v)
		maxVal = max(maxVal, v)
	}
	res := make([]int, 0, maxVal-minVal)
	for i := minVal + 1; i < maxVal; i++ {
		if !everyInteger[i] {
			res = append(res, i)
		}
	}
	return res
}

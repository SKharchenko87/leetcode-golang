package p2200

func findKDistantIndices(nums []int, key int, k int) []int {
	l := len(nums)
	index := 0
	res := make([]int, 0, l)
	for i := 0; i < l; i++ {
		if nums[i] == key {
			for j := max(index, i-k); j < min(l, i+k+1); j++ {
				res = append(res, j)
			}
			index = min(l, i+k+1)
		}
	}
	return res
}

func findKDistantIndices0(nums []int, key int, k int) []int {
	l := len(nums)
	indexes := make([]bool, l)
	for i := 0; i < l; i++ {
		if nums[i] == key {
			for j := max(0, i-k); j < min(l, i+k+1); j++ {
				indexes[j] = true
			}
		}
	}
	res := make([]int, 0, l)
	for i := 0; i < l; i++ {
		if indexes[i] {
			res = append(res, i)
		}
	}
	return res
}

package p3375

func minOperations(nums []int, k int) int {
	freq := [101]int{}
	res := 0
	for _, n := range nums {
		if n < k {
			return -1
		}
		freq[n]++
		if n > k && freq[n] == 1 {
			res++
		}
	}
	return res
}

package p3512

func minOperations(nums []int, k int) int {
	d := 0
	for _, v := range nums {
		d = (d + v) % k
	}
	return d
}

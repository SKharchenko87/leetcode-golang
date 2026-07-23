package p3513

func uniqueXorTriplets(nums []int) int {
	n := len(nums)
	if n < 3 {
		return n
	}
	x := 1
	for n/x > 0 {
		x <<= 1
	}
	return x
}

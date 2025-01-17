package p2425

func xorAllNums(nums1 []int, nums2 []int) int {
	l1, l2 := len(nums1), len(nums2)
	var xor1, xor2 int
	if l2&1 == 1 {
		for i := 0; i < l1; i++ {
			xor1 ^= nums1[i]
		}
	}
	if l1&1 == 1 {
		for i := 0; i < l2; i++ {
			xor2 ^= nums2[i]
		}
	}
	return xor1 ^ xor2
}

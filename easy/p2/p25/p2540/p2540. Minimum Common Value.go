package p2540

func getCommon(nums1 []int, nums2 []int) int {
	l1, l2 := len(nums1), len(nums2)
	i1, i2 := 0, 0
	for i1 < l1 && i2 < l2 {
		if nums1[i1] == nums2[i2] {
			return nums1[i1]
		}
		m := max(nums1[i1], nums2[i2])
		for ; i1 < l1 && nums1[i1] < m; i1++ {
		}
		for ; i2 < l2 && nums2[i2] < m; i2++ {
		}
	}
	return -1
}

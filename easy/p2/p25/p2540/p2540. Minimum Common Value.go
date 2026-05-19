package p2540

import "slices"

func getCommon(nums1 []int, nums2 []int) int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	for i := 0; i < len(nums1); i++ {
		if _, ok := slices.BinarySearch(nums2, nums1[i]); ok {
			return nums1[i]
		}
	}
	return -1
}

func getCommon1(nums1 []int, nums2 []int) int {
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			return nums1[i]
		}
		if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return -1
}

func getCommon0(nums1 []int, nums2 []int) int {
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

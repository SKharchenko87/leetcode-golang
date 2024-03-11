package p0349

import "sort"

func intersection(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	len1, len2 := len(nums1), len(nums2)
	i, j := 0, 0
	res := []int{}
	for i < len1 && j < len2 {
		tmp := nums1[i]
		for ; j < len2 && nums2[j] < tmp; j++ {

		}
		if j < len2 && tmp == nums2[j] {
			res = append(res, tmp)
		}
		for ; i < len1 && nums1[i] == tmp; i++ {

		}
	}
	return res
}

package p2040

import "sort"

//TLE
func kthSmallestProduct(nums1 []int, nums2 []int, k int64) int64 {
	l1, l2 := len(nums1), len(nums2)
	res := make([]int64, 0, l1*l2)
	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			res = append(res, int64(nums1[i])*int64(nums2[j]))
		}
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	return res[k-1]
}

package p0350

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	tmp := map[int]int{}
	res := []int{}
	for i := 0; i < len(nums1); i++ {
		tmp[nums1[i]]++
	}
	for i := 0; i < len(nums2); i++ {
		if tmp[nums2[i]] > 0 {
			res = append(res, nums2[i])
			tmp[nums2[i]]--
		}
	}
	return res
}

func intersect1(nums1 []int, nums2 []int) []int {
	tmp := [1001]int{}
	res := []int{}
	for i := 0; i < len(nums1); i++ {
		tmp[nums1[i]]++
	}
	for i := 0; i < len(nums2); i++ {
		if tmp[nums2[i]] > 0 {
			res = append(res, nums2[i])
			tmp[nums2[i]]--
		}
	}
	return res
}

func intersect0(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	res := []int{}
	j := 0
	for i := 0; i < len(nums1); i++ {
		for ; j < len(nums2); j++ {
			if nums1[i] < nums2[j] {
				break
			}
			if nums1[i] == nums2[j] {
				res = append(res, nums1[i])
				j++
				break
			}
		}
	}
	return res
}

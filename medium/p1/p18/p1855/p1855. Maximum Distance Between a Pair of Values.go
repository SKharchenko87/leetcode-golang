package p1855

func maxDistance(nums1 []int, nums2 []int) int {
	n1, n2 := len(nums1), len(nums2)
	i, j := n1-1, n2-1
	res := 0
	for ; j >= 0 && i >= 0; j-- {
		for ; i >= 0 && nums1[i] <= nums2[j]; i-- {
			if i <= j {
				res = max(res, j-i)
			}
		}
	}
	return res
}

//func maxDistance(nums1 []int, nums2 []int) int {
//	res := 0
//	i := len(nums1) - 1
//	for j, n2 := range slices.Backward(nums2) {
//		for ; i >= 0 && nums1[i] <= n2; i-- {
//			if i <= j {
//				res = max(res, j-i)
//			}
//		}
//		if i < 0 {
//			break
//		}
//	}
//	return res
//}

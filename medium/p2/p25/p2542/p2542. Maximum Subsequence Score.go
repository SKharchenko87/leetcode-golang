package p2542

import (
	"fmt"
	"math"
	"sort"
)

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	sort.Slice(nums1, func(i, j int) bool {
		return nums2[i] > nums2[j]
	})
	sort.Ints(nums2)
	reverseSlice(nums2)
	fmt.Println(nums1, nums2)
	l := len(nums1)
	minM := math.MaxInt
	sum := 0
	for i := k; i > 0; i-- {
		if minM > nums2[l-i] {
			minM = nums2[l-i]
		}
		sum += nums1[l-i]
	}
	return int64(minM) * int64(sum)
}

func reverseSlice(s []int) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		s[i], s[length-i-1] = s[length-i-1], s[i]
	}
}

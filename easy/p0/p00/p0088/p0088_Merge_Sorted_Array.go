package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	i1 := 0
	i2 := 0
	for {
		if i1 < m {
			if nums1[i1] <= nums2[i2] {
				i1++
			} else {
				//сдвигаем
				for k := m; k > i1; k-- {
					nums1[k] = nums1[k-1]
				}
				//fmt.Println(nums1)
				//заменяем
				nums1[i1] = nums2[i2]
				i2++
				m++
				fmt.Println(nums1)
			}
		} else {
			nums1[i1] = nums2[i2]
			i1++
			i2++
			m++
		}
		if i1 == len(nums1) || m == len(nums1) {
			break
		}
	}
}

func main() {
	nums1 := []int{4, 5, 6, 0, 0, 0}
	m := 3
	nums2 := []int{1, 2, 3}
	n := 3
	//Output: [1,2,2,3,5,6]
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}

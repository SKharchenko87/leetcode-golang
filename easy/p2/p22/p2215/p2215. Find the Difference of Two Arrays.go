package main

import "fmt"

func intsToMap(nums []int, c chan map[int]int) {
	m := make(map[int]int, len(nums))
	for _, v := range nums {
		m[v]++
	}
	c <- m
}

func mapToInts(m map[int]int, c chan []int) {
	i := 0
	arr := make([]int, len(m))
	for k, _ := range m {
		arr[i] = k
		i++
	}
	c <- arr
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	var ch1 = make(chan map[int]int)
	var ch2 = make(chan map[int]int)
	go intsToMap(nums1, ch1)
	go intsToMap(nums2, ch2)
	m1 := <-ch1
	m2 := <-ch2

	for k, _ := range m1 {
		if _, ok := m2[k]; ok {
			delete(m1, k)
			delete(m2, k)
		}
	}

	var chr1 = make(chan []int)
	var chr2 = make(chan []int)
	go mapToInts(m1, chr1)
	go mapToInts(m2, chr2)
	return [][]int{<-chr1, <-chr2}
}

func main() {
	nums1, nums2 := []int{1, 2, 3}, []int{2, 4, 6}
	fmt.Println(findDifference(nums1, nums2))
}

//func findDifference(nums1 []int, nums2 []int) [][]int {
//	m1 := make(map[int]int, len(nums1))
//	m2 := make(map[int]int, len(nums2))
//	for _, v := range nums1 {
//		m1[v]++
//	}
//	for _, v := range nums2 {
//		m2[v]++
//	}
//	for k, _ := range m1 {
//		if _, ok := m2[k]; ok {
//			delete(m1, k)
//			delete(m2, k)
//		}
//	}
//	r1 := make([]int, len(m1))
//	i := 0
//	for k, _ := range m1 {
//		r1[i] = k
//		i++
//	}
//	r2 := make([]int, len(m2))
//	i = 0
//	for k, _ := range m2 {
//		r2[i] = k
//		i++
//	}
//	return [][]int{r1, r2}
//
//}

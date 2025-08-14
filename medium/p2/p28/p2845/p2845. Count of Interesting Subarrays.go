package p2845

import "slices"

func countInterestingSubarrays(nums []int, modulo int, k int) int64 {
	
	res := int64(0)
	currentCount := 0
	m := map[int]int{}
	m[0] = 1
	for _, num := range nums {
		if num%modulo == k {
			currentCount++
		}
		counts := currentCount % modulo
		res += int64(m[(modulo+currentCount-k)%modulo])
		m[counts]++
	}
	return res
}

func countInterestingSubarrays1(nums []int, modulo int, k int) int64 {
	//l := len(nums)

	currentCount := 0
	m := map[int][]int{}
	for i, num := range nums {
		if num%modulo == k {
			currentCount++
		}
		counts := currentCount % modulo
		if _, ok := m[counts]; !ok {
			m[counts] = []int{}
		}
		m[counts] = append(m[counts], i)
	}
	res := int64(0)
	for key, ints := range m {
		for _, i := range ints {
			index, _ := slices.BinarySearch(m[(k+key)%modulo], i+1)
			res += int64(len(m[(k+key)%modulo]) - index)
		}
	}
	res += int64(len(m[(k)]))
	return res
}

func countInterestingSubarrays0(nums []int, modulo int, k int) int64 {
	l := len(nums)
	counts := make([]int, l)
	currentCount := 0
	m := map[int][]int{}
	for i, num := range nums {
		if num%modulo == k {
			currentCount++
		}
		counts[i] = currentCount % modulo
		if _, ok := m[counts[i]]; !ok {
			m[counts[i]] = []int{}
		}
		m[counts[i]] = append(m[counts[i]], i)
	}
	res := int64(0)
	for key, ints := range m {
		for _, i := range ints {
			index, _ := slices.BinarySearch(m[(k+key)%modulo], i+1)
			res += int64(len(m[(k+key)%modulo]) - index)
		}
	}
	res += int64(len(m[(k)]))
	return res
}

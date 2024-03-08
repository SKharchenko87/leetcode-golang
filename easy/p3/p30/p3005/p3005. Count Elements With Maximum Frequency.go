package p3005

import "sort"

func maxFrequencyElements(nums []int) int {
	m := map[int]int{}
	maxFrequency := 0
	count := 0
	for _, n := range nums {
		m[n]++
		if m[n] == maxFrequency {
			count++
		} else if m[n] > maxFrequency {
			count = 1
			maxFrequency = m[n]
		}
	}
	return count * maxFrequency
}

func maxFrequencyElements0(nums []int) int {
	l := len(nums)
	m := map[int]int{}
	sort.Ints(nums)
	count := 1
	maxFrequency := 0
	for i := 1; i < l; i++ {
		if nums[i-1] == nums[i] {
			count++
		} else {
			m[count]++
			if maxFrequency < count {
				maxFrequency = count
			}
			count = 1
		}
	}
	m[count]++
	if maxFrequency < count {
		maxFrequency = count
	}
	return m[maxFrequency] * maxFrequency
}

package p1636

import "sort"

func frequencySort(nums []int) []int {
	freq := make([]int, 201)
	for _, num := range nums {
		freq[num+100]++
	}
	sort.SliceStable(nums, func(i, j int) bool {
		if freq[nums[i]+100] == freq[nums[j]+100] {
			return nums[i] > nums[j]
		}
		return freq[nums[i]+100] < freq[nums[j]+100]
	})
	return nums
}

func frequencySort2(nums []int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	sort.SliceStable(nums, func(i, j int) bool {
		if freq[nums[i]] == freq[nums[j]] {
			return nums[i] > nums[j]
		}
		return freq[nums[i]] < freq[nums[j]]
	})
	return nums
}

func frequencySort1(nums []int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	sort.Slice(nums, func(i, j int) bool {
		if freq[nums[i]] < freq[nums[j]] {
			return true
		} else if freq[nums[i]] == freq[nums[j]] {
			return nums[i] > nums[j]
		}
		return false
	})
	return nums
}

type frequency struct {
	nums *[]int
	freq *map[int]int
}

func (f frequency) Len() int {
	return len(*f.nums)
}

func (f frequency) Less(i, j int) bool {
	if (*f.freq)[(*f.nums)[i]] < (*f.freq)[(*f.nums)[j]] {
		return true
	} else if (*f.freq)[(*f.nums)[i]] == (*f.freq)[(*f.nums)[j]] {
		return (*f.nums)[i] > (*f.nums)[j]
	}
	return false
}

func (f frequency) Swap(i, j int) {
	(*f.nums)[i], (*f.nums)[j] = (*f.nums)[j], (*f.nums)[i]
}

func frequencySort0(nums []int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	f := frequency{nums: &nums, freq: &freq}
	sort.Sort(f)
	return nums
}

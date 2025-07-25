package p3487

import "slices"

func maxSum(nums []int) int {
	set := [101]bool{}
	maxNonPositiveElement := -101
	result := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			if !set[nums[i]] {
				result += nums[i]
				set[nums[i]] = true
			}
		} else if maxNonPositiveElement < nums[i] {
			maxNonPositiveElement = nums[i]
		}
	}
	if result == 0 {
		return maxNonPositiveElement
	}
	return result
}

func maxSum2(nums []int) int {
	slices.Sort(nums)
	l := len(nums)
	result := nums[l-1]
	for i := l - 2; i >= 0 && nums[i] > 0; i-- {
		if nums[i] != nums[i+1] {
			result += nums[i]
		}
	}
	return result
}

func maxSum1(nums []int) int {
	set := make(map[int]struct{}, 100)
	maxNonPositiveElement := -101
	result := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			if _, ok := set[nums[i]]; !ok {
				result += nums[i]
				set[nums[i]] = struct{}{}
			}
		} else if maxNonPositiveElement < nums[i] {
			maxNonPositiveElement = nums[i]
		}
	}
	if result == 0 {
		return maxNonPositiveElement
	}
	return result
}

func maxSum0(nums []int) int {
	freq := [201]bool{}
	for _, num := range nums {
		freq[num+100] = true
	}
	res := 0
	i := 200
	for ; i > 100; i-- {
		if freq[i] {
			res += i - 100
		}
	}
	if res == 0 {
		for ; !freq[i]; i-- {
		}
		res += i - 100
	}
	return res
}

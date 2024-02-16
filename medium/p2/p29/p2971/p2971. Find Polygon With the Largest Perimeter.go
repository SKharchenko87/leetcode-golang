package p2971

import "sort"

// 1, 1, 2, 3,  5, 12, 50
// 1, 2, 4, 7, 12, 24, 74 - sum
//-1,-1, 4, 7, 12, 12, 12

func largestPerimeter(nums []int) int64 {
	sort.Ints(nums)
	var res, sum int64 = -1, int64(nums[0] + nums[1])
	for i := 2; i < len(nums); i++ {
		if sum > int64(nums[i]) {
			res = sum + int64(nums[i])
		}
		sum += int64(nums[i])
	}
	return res
}

func largestPerimeter2(nums []int) int64 {
	l := len(nums)
	sort.Ints(nums)
	var sum int64
	for _, n := range nums {
		sum += int64(n)
	}
	for i := l - 1; i >= 0; i-- {
		sum -= int64(nums[i])
		if sum > int64(nums[i]) {
			return sum + int64(nums[i])
		}
	}
	return -1
}

func largestPerimeter1(nums []int) int64 {
	var sum int64
	sort.Ints(nums)
	for j := len(nums) - 1; j >= 0; j-- {
		sum = sum + int64(nums[j])
	}

	i := len(nums) - 1
	for ; i >= 2; i-- {
		v := nums[i]
		sum = sum - int64(v)
		if sum > int64(v) {
			sum = sum + int64(v)
			break
		}
	}
	if i <= 1 {
		return -1
	}
	return sum
}

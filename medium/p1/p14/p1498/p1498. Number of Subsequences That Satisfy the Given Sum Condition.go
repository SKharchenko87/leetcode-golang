package p1498

import (
	"slices"
	"sort"
)

const MOD int = 1e9 + 7
const SIZE int = 1e5

var pow2mod = make([]int, SIZE)

func init() {
	pow2mod[0] = 1
	for i := 1; i < SIZE; i++ {
		pow2mod[i] = pow2mod[i-1] * 2 % MOD
	}
}

func numSubseq(nums []int, target int) int {
	sort.Ints(nums)
	result := 0

	right, _ := slices.BinarySearch(nums, target-nums[0]+1)
	right = min(len(nums)-1, right)
	left := 0
	for left <= right {
		if nums[left]+nums[right] > target {
			right--
		} else {
			result = (result + pow2mod[right-left]) % MOD
			left++
		}
	}
	return result
}

func numSubseq0(nums []int, target int) int {
	sort.Ints(nums)
	result := 0
	for i := 0; i < len(nums) && 2*nums[i] <= target; i++ {
		index, _ := slices.BinarySearch(nums, target-nums[i]+1)
		p := index - i
		result += powerWithMOD(2, p-1, MOD)
		result %= MOD
	}
	return result
}

func powerWithMOD[T int](x, y, mod T) T {
	var res T = 1
	for y != 0 {
		if y%2 == 1 {
			res = res * x % mod
		}
		x = x * x % mod
		y /= 2
	}
	return res
}

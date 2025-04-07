package p0416

import (
	"maps"
	"sort"
)

func canPartition(nums []int) bool {
	l := len(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	// Проверяем что массив можно разделить на две части
	if sum%2 == 1 {
		return false
	}
	// Проверяем что максимальное число меньше суммы одной части(половины общей суммы)
	partSum := sum / 2
	if nums[l-1] > partSum {
		return false
	}
	maxVal := l * 200
	dp := make([]bool, maxVal)
	dp[0] = true
	for i := 0; i < l; i++ {
		for j := len(dp) - 1; j >= 0; j-- {
			if dp[j] {
				if nums[i]+j < maxVal && !dp[nums[i]+j] {
					dp[nums[i]+j] = true
					if nums[i]+j == partSum {
						return true
					}
				}
			}
		}
	}
	return false
}

func canPartition2(nums []int) bool {
	l := len(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	// Проверяем что массив можно разделить на две части
	if sum%2 == 1 {
		return false
	}
	// Проверяем что максимальное число меньше суммы одной части(половины общей суммы)
	partSum := sum / 2
	if nums[l-1] > partSum {
		return false
	}
	dp := make(map[int]bool, l*l)
	dp[0] = true
	for _, num := range nums {
		newDP := maps.Clone(dp)
		for key := range dp {
			newDP[key+num] = true
			if newDP[partSum] {
				return true
			}
		}
		dp = newDP
	}
	return false
}

func canPartition0(nums []int) bool {
	l := len(nums)
	sort.Ints(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	// Проверяем что массив можно разделить на две части
	if sum%2 == 1 {
		return false
	}
	// Проверяем что максимальное число меньше суммы одной части(половины общей суммы)
	partSum := sum / 2
	//if nums[l-1] > partSum {
	//	return false
	//}
	dp := make(map[int]bool, l*l)
	dp[0] = true
	for _, num := range nums {
		newDP := maps.Clone(dp)
		for key := range dp {
			newDP[key+num] = true
			if newDP[partSum] {
				return true
			}
		}
		dp = newDP
	}
	return false
}

func canPartition1(nums []int) bool {
	l := len(nums)
	sort.Ints(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	// Проверяем что массив можно разделить на две части
	if sum%2 == 1 {
		return false
	}
	// Проверяем что максимальное число меньше суммы одной части(половины общей суммы)
	partSum := sum / 2
	if nums[l-1] > partSum {
		return false
	}
	dp := make(map[int]bool, l*l)
	dp[0] = true
	for _, num := range nums {
		newDP := maps.Clone(dp)
		for key := range dp {
			newDP[key+num] = true
			if newDP[partSum] {
				return true
			}
		}
		dp = newDP
	}
	return false
}

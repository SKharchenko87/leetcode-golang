package p0322

import (
	"math"
	"sort"
)

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}

	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			dp[j] = min(dp[j], dp[j-coin]+1)
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func coinChange1(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}
	sort.Ints(coins)
	for i := 0; i <= amount; i++ {
		for j := 0; j < len(coins) && coins[j] <= i; j++ {
			dp[i] = min(dp[i-coins[j]]+1, dp[i])
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func coinChange0(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}
	sort.Ints(coins)
	for j := 0; j < len(coins) && coins[j] <= amount; j++ {
		dp[coins[j]] = 1
	}
	for i := 0; i <= amount; i++ {
		for j := 0; j < len(coins) && coins[j] <= i; j++ {
			dp[i] = min(dp[i-coins[j]]+1, dp[i])
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

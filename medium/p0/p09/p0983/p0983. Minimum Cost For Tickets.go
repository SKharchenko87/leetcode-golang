package p0983

import "slices"

func mincostTickets(days []int, costs []int) int {
	dp := make([]int, len(days)+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {

		if dp[i] == 0 || dp[i] > dp[i-1]+costs[0] {
			dp[i] = dp[i-1] + costs[0]
		}
		index, _ := slices.BinarySearch(days, days[i-1]+7)
		if dp[index] == 0 || dp[index] > dp[i-1]+costs[1] {
			dp[index] = dp[i-1] + costs[1]
		}

		index, _ = slices.BinarySearch(days, days[i-1]+30)
		if dp[index] == 0 || dp[index] > dp[i-1]+costs[2] {
			dp[index] = dp[i-1] + costs[2]
		}
	}
	return dp[len(dp)-1]
}

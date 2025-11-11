package p0474

import "strings"

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for k := 0; k < len(strs); k++ {
		zero := strings.Count(strs[k], "0")
		one := len(strs[k]) - zero
		for i := m; i >= zero; i-- {
			for j := n; j >= one; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zero][j-one]+1)
			}
		}
	}
	return dp[m][n]
}

func findMaxForm1(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	var one, zero int
	for k := 0; k < len(strs); k++ {
		one, zero = 0, 0
		for _, ch := range strs[k] {
			if ch == '0' {
				zero++
			} else {
				one++
			}
		}

		for i := m; i >= zero; i-- {
			for j := n; j >= one; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zero][j-one]+1)
			}
		}
	}
	return dp[m][n]
}

func findMaxForm0(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	l := len(strs)

	one := make([]int, l)
	zero := make([]int, l)
	for k := 0; k < l; k++ {
		for _, ch := range strs[k] {
			if ch == '0' {
				zero[k]++
			} else {
				one[k]++
			}
		}
		for i := m; i >= zero[k]; i-- {
			for j := n; j >= one[k]; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zero[k]][j-one[k]]+1)
			}
		}
	}
	return dp[m][n]
}

/*

"10", "0001", "111001", "1", "0"
1     2       4        8     16


10 10 10 111 11


m=1
n=5

*/

package p3665

const mod int = 1e9 + 7

/*ToDo если не использовать interface то должно быть быстрее*/
func uniquePaths(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	dp := make([][]interface{}, n)
	for i := range dp {
		dp[i] = make([]interface{}, m)
	}
	dp[0][0] = 1
	var ok bool
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var up, left int
			if i > 0 {
				if up, ok = dp[i-1][j].(int); !ok {
					up = dp[i-1][j].([2]int)[0]
				}
			}

			if j > 0 {
				if left, ok = dp[i][j-1].(int); !ok {
					left = dp[i][j-1].([2]int)[1]
				}
			}
			if grid[i][j] == 0 {
				if dp[i][j] == nil {
					dp[i][j] = 0
				}
				dp[i][j] = (dp[i][j].(int) + up + left) % mod
			} else {
				dp[i][j] = [2]int{left % mod, up % mod}
			}
		}
	}
	return dp[n-1][m-1].(int)
}

func uniquePaths0(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	dp := make([][]interface{}, n)
	for i := range dp {
		dp[i] = make([]interface{}, m)
	}
	dp[0][0] = 1
	var x int
	var ok bool
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 0 {
				x = 0
				if dp[i][j] == nil {
					dp[i][j] = 0
				}
				if i > 0 {
					if x, ok = dp[i-1][j].(int); !ok {
						x = dp[i-1][j].([2]int)[0]
					}
				}
				dp[i][j] = (dp[i][j].(int) + x) % mod
				x = 0
				if j > 0 {
					if x, ok = dp[i][j-1].(int); !ok {
						x = dp[i][j-1].([2]int)[1]
					}
				}
				dp[i][j] = (dp[i][j].(int) + x) % mod
			} else { // mirror
				if dp[i][j] == nil {
					dp[i][j] = [2]int{}
				}
				t := dp[i][j].([2]int)
				if i > 0 {
					x = 0
					if x, ok = dp[i-1][j].(int); !ok {
						x = dp[i-1][j].([2]int)[0]
					}
					t[1] = (t[1] + x) % mod
				}
				if j > 0 {
					x = 0
					if x, ok = dp[i][j-1].(int); !ok {
						x = (dp[i][j-1].([2]int)[1]) % mod
					}
					t[0] = (t[0] + x) % mod
				}
				dp[i][j] = t
			}
		}
	}
	return dp[n-1][m-1].(int)
}

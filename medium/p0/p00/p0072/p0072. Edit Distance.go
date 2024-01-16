package p0072

import "fmt"

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	//if m == 0 && n == 0 {
	//	return 0
	//} else if m == 0 {
	//	return n
	//} else if n == 0 {
	//	return m
	//}
	m++
	n++
	d := make([][]int, m)
	for i := 0; i < m; i++ {
		d[i] = make([]int, n)
		d[i][0] = i
	}
	for j := 1; j < n; j++ {
		d[0][j] = j
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if word1[i-1] == word2[j-1] {
				d[i][j] = d[i-1][j-1]
			} else {
				d[i][j] = min(d[i][j-1], d[i-1][j], d[i-1][j-1]) + 1
			}
		}
	}

	fmt.Println(d)
	return d[m-1][n-1]
}

func minDistance2(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}

	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if word1[r] == word2[c] {
				dp[r+1][c+1] = dp[r][c]
			} else {
				dp[r+1][c+1] = min(dp[r][c], min(dp[r+1][c], dp[r][c+1])) + 1
			}
		}
	}
	fmt.Println(dp)
	return dp[m][n]
}

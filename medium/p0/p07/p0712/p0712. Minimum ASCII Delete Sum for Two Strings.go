/*
	d e l e t e

l  . . * . . .
e  . * . * . *
e  . * . * . *
t  . . . . * .

l
le
e
t

lt
ee
et
let
lee
eet
*/
package p0712

func minimumDeleteSum(s1 string, s2 string) int {
	totalSum := 0
	l1, l2 := len(s1), len(s2)
	for i := 0; i < l1; i++ {
		totalSum += int(s1[i])
	}
	for i := 0; i < l2; i++ {
		totalSum += int(s2[i])
	}
	dp := make([][]int, l1+1)
	dp[0] = make([]int, l2+1)
	for i := 1; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
	}

	var cross int
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if s1[i-1] == s2[j-1] {
				cross = int(s1[i-1])
			} else {
				cross = 0
			}
			dp[i][j] = max(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]+cross)
		}
	}
	return totalSum - 2*dp[l1][l2]
}

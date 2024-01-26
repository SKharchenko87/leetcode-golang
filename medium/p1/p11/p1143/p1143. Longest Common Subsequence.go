package p1143

/*
	   o  x  c  p  q  r  s  v  w  f
	0  0  0  0  0  0  0  0  0  0  0

s 0  0  0  0  0  0  0  1  0  0  0
h 0  0  0  0  0  0  0  1  1  0  0
m 0  0  0  0  0  0  0  1  1  1  0
t 0  0  0  0  0  0  0  1  1  1  1
u 0  0  0  0  0  0  0  1  1  1  1
l 0  0  0  0  0  0  0  1  1  1  1
q 0  0  0  0  0  1  0  1  1  1  1
r 0  0  0  0  0  1  2  2  2  2  2
y 0  0  0  0  0  1  2  2  2  2  2
p 0  0  0  0  1  1  2  2  2  2  2
y 0  0  0  0  1  1  2  2  2  2  2
*/
func longestCommonSubsequence(text1 string, text2 string) int {
	l1, l2 := len(text1)+1, len(text2)+1
	dp := make([][]int, l1)
	dp[0] = make([]int, l2)
	for i := 1; i < l1; i++ {
		dp[i] = make([]int, l2)
		for j := 1; j < l2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[l1-1][l2-1]
}

package p2140

func mostPoints(questions [][]int) int64 {
	l := len(questions)
	dp := make([]int64, l+1)
	currentPoint := int64(questions[0][0])
	nextIndex := questions[0][1] + 1
	if nextIndex > l {
		nextIndex = l
	}
	if dp[nextIndex] < currentPoint {
		dp[nextIndex] = currentPoint
	}
	for i := 1; i < l; i++ {
		dp[i] = max(dp[i-1], dp[i])
		currentPoint = int64(questions[i][0]) + dp[i]
		nextIndex = i + questions[i][1] + 1
		if nextIndex > l {
			nextIndex = l
		}
		if dp[nextIndex] < currentPoint {
			dp[nextIndex] = currentPoint
		}
	}
	return dp[l]
}

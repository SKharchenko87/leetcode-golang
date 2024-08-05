package p1105

func minHeightShelves(books [][]int, shelfWidth int) int {
	l := len(books)
	dp := make([]int, l+1)
	dp[0] = 0
	dp[1] = books[0][1]

	for i := 2; i < l+1; i++ {
		j := i - 1
		remainingShelfWidth := shelfWidth - books[j][0]
		maxHeight := books[j][1]
		dp[i] = dp[j] + books[j][1]
		for j > 0 && remainingShelfWidth-books[j-1][0] >= 0 {
			maxHeight = max(maxHeight, books[j-1][1])
			remainingShelfWidth -= books[j-1][0]
			dp[i] = min(dp[i], maxHeight+dp[j-1])
			j--
		}
	}
	return dp[l]
}

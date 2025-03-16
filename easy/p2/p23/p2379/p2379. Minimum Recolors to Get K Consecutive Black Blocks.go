package p2379

func minimumRecolors(blocks string, k int) int {
	n := len(blocks)
	sum := 0
	i := 0
	for ; i < k; i++ {
		if blocks[i] == 'B' {
			sum += 1
		}
	}
	maxSum := sum
	for ; i < n; i++ {
		if blocks[i] == 'B' {
			sum += 1
		}
		if blocks[i-k] == 'B' {
			sum -= 1
		}
		maxSum = max(maxSum, sum)
	}
	return k - maxSum
}

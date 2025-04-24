package p2145

func numberOfArrays(differences []int, lower int, upper int) int {
	minSum, maxSum := 0, 0
	sum := 0
	for _, difference := range differences {
		sum += difference
		maxSum = max(maxSum, sum)
		minSum = min(minSum, sum)
	}
	d := upper - lower - (maxSum - minSum) + 1
	if d >= 0 {
		return d
	}
	return 0
}

package p1335

func minDifficulty(jobDifficulty []int, d int) int {
	n := len(jobDifficulty)
	sum := 0
	if n < d {
		return -1
	} else if n == d {
		for _, v := range jobDifficulty {
			sum += v
		}
		return sum // Efficiently sum the difficulties
	}

	// dp[j]: minimum difficulty of first (j+1) jobs in (i+1) days
	dp := make([]int, n)
	dp[0] = jobDifficulty[0]

	// Initialize dp with maximum difficulty encountered so far
	for i := 1; i < n; i++ {
		dp[i] = max(jobDifficulty[i], dp[i-1])
	}

	dpPrev := make([]int, n)
	//stack := []int{}  // Use a list for the decreasing stack

	// Dynamic Programming with decreasing stack optimization
	for i := 1; i < d; i++ {
		dp, dpPrev = dpPrev, dp // Swap dp arrays concisely
		stack := []int{}        // Clear the stack efficiently

		for j := i; j < n; j++ {
			dp[j] = jobDifficulty[j] + dpPrev[j-1]

			// Process the decreasing stack
			for len(stack) > 0 && jobDifficulty[stack[len(stack)-1]] <= jobDifficulty[j] {
				lastIdx := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				dp[j] = min(dp[j], dp[lastIdx]+jobDifficulty[j]-jobDifficulty[lastIdx])
			}

			if len(stack) > 0 {
				dp[j] = min(dp[j], dp[stack[len(stack)-1]])
			}
			stack = append(stack, j) // Append for clarity and efficiency
		}
	}
	return dp[n-1]
}

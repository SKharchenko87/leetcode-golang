package p0440

func findKthNumber(n int, k int) int {
	curr := 1
	k--

	for k > 0 {
		steps := calcSteps(n, curr, curr+1)
		if steps <= k {
			curr++
			k -= steps
		} else {
			curr *= 10
			k--
		}
	}

	return curr
}

func calcSteps(n, curr, next int) int {
	steps := 0
	for curr <= n {
		steps += min(n+1, next) - curr
		curr *= 10
		next *= 10
	}

	return steps
}

func findKthNumber0(n int, k int) int {
	var res int
	//res = make([]int, n)
	index := 0
	pow10 := []int{0, 1, 10, 100, 1000, 10_000, 100_000, 1_000_000, 10_000_000, 100_000_000, 1_000_000_000, 10_000_000_000}
	var dfs func(level int)
	dfs = func(level int) {

		for index <= k && index < n && pow10[level] < (pow10[level-1]+1)*10 && pow10[level] <= n {
			if index == k-1 {
				res = pow10[level]
				index++
				return
				//return pow10[level]
			}
			//res[index] = pow10[level]
			index++
			dfs(level + 1)
			pow10[level]++
		}
	}
	dfs(1)
	return res
}

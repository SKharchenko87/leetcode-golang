package p1718

func constructDistancedSequence(n int) []int {
	result := make([]int, 2*n-1)
	dfs(n, 0, 0, &result)
	return result
}

func dfs(n, index, usedBitMap int, result *[]int) bool {
	if index == len(*result) {
		return true
	}
	if (*result)[index] > 0 {
		return dfs(n, index+1, usedBitMap, result)
	}

	for num := n; num > 0; num-- {
		if usedBitMap&(1<<num) != 0 {
			continue
		}
		if num > 1 && (index+num >= len(*result) || (*result)[index+num] > 0) {
			continue
		}

		(*result)[index] = num
		if num > 1 {
			(*result)[index+num] = num
		}

		if dfs(n, index+1, usedBitMap^(1<<num), result) {
			return true
		}
		(*result)[index] = 0
		if num > 1 {
			(*result)[index+num] = 0
		}
	}
	return false
}

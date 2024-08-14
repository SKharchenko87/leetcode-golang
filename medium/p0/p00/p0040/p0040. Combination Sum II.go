package p0040

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	l := len(candidates)
	result := make([][]int, 0)
	var dfs func(start, target int, path []int)
	dfs = func(start, target int, path []int) {
		if target < 0 {
			return // backtracking
		}
		if target == 0 {
			newPath := make([]int, len(path))
			copy(newPath, path)
			result = append(result, newPath)
			return
		}
		for i := start; i < l; i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			dfs(i+1, target-candidates[i], append(path, candidates[i]))
		}
	}
	dfs(0, target, []int{})
	return result
}

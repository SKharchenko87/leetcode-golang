package p0216

func dfs(k, n int, res *[][]int, tmp []int, start int) {
	if k == 0 && n == 0 {
		newArr := make([]int, len(tmp))
		copy(newArr, tmp)
		*res = append(*res, newArr)
	}
	if n < 0 || k == 0 {
		return
	}
	for i := start; i <= 10-k; i++ {
		dfs(k-1, n-i, res, append(tmp, i), i+1)
	}
}

func combinationSum3(k int, n int) [][]int {
	res := [][]int{}
	dfs(k, n, &res, []int{}, 1)
	return res
}

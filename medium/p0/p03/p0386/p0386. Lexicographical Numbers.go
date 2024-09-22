package p0386

import (
	"slices"
	"strconv"
)

func lexicalOrder(n int) (res []int) {
	res = make([]int, n)
	index := 0
	pow10 := []int{1, 1, 10, 100, 1000, 10000, 100000}
	var dfs func(level int)
	dfs = func(level int) {
		for index < n && pow10[level] < (pow10[level-1]+1)*10 && pow10[level] <= n {
			res[index] = pow10[level]
			index++
			dfs(level + 1)
			pow10[level]++
		}
	}
	dfs(1)
	return
}

func lexicalOrder2(n int) []int {
	res := make([]int, n)
	index := 0
	pow10 := []int{1, 10, 100, 1000, 10000, 100000}
	var dfs func(level int)
	dfs = func(level int) {
		for pow10[level+1] < (pow10[level]+1)*10 && pow10[level+1] <= n {
			res[index] = pow10[level+1]
			index++
			dfs(level + 1)
			pow10[level+1]++
		}
	}

	for i := 1; i < 10; i++ {
		res[index] = i
		index++
		dfs(0)
		if index == n {
			return res
		}
		pow10[0]++
	}
	return res
}

func lexicalOrder1(n int) []int {
	res := make([]int, n)
	index := 0
	i10 := 10
	i100 := 100
	i1000 := 1000
	i10000 := 10000
	i100000 := 100000
	for i := 1; i < 10; i++ {
		res[index] = i
		index++
		for ; i10 < (i+1)*10 && i10 <= n; i10++ {
			res[index] = i10
			index++
			for ; i100 < (i10+1)*10 && i100 <= n; i100++ {
				res[index] = i100
				index++
				for ; i1000 < (i100+1)*10 && i1000 <= n; i1000++ {
					res[index] = i1000
					index++
					for ; i10000 < (i1000+1)*10 && i10000 <= n; i10000++ {
						res[index] = i10000
						index++
						for ; i100000 < (i10000+1)*10 && i100000 <= n; i100000++ {
							res[index] = i100000
							index++
						}
					}
				}
			}
		}

		if index == n {
			return res
		}
	}
	return res
}

func lexicalOrder0(n int) []int {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = i + 1
	}
	slices.SortFunc(res, func(a, b int) int {
		if strconv.Itoa(a) < strconv.Itoa(b) {
			return -1
		} else if strconv.Itoa(a) == strconv.Itoa(b) {
			return 0
		}
		return 1
	})
	return res
}

func run(n int) bool {
	x := slices.Compare(lexicalOrder(n), lexicalOrder0(n))
	if x == 0 {
		return true
	}
	return false
}

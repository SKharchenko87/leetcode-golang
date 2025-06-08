package p0386

import (
	"slices"
	"strconv"
)

func lexicalOrder(n int) []int {
	res := make([]int, n)
	index := 0
	is := []int{1, 10, 100, 1000, 10000}
	for i := 1; i < 10 && i <= n; i++ {
		res[index] = i
		index++
		for ; is[1] < (i+1)*10 && is[1] <= n; is[1]++ {
			res[index] = is[1]
			index++
			for ; is[2] < (is[1]+1)*10 && is[2] <= n; is[2]++ {
				res[index] = is[2]
				index++
				for ; is[3] < (is[2]+1)*10 && is[3] <= n; is[3]++ {
					res[index] = is[3]
					index++
					for ; is[4] < (is[3]+1)*10 && is[4] <= n; is[4]++ {
						res[index] = is[4]
						index++
					}
				}
			}
		}
	}
	return res
}

func lexicalOrder4(n int) []int {
	res := make([]int, 0, n)
	var dfs func(prefix int)
	dfs = func(prefix int) {
		if prefix > n {
			return
		}
		res = append(res, prefix)
		if prefix*10 <= n {
			dfs(prefix * 10)
		}
		if prefix+1 <= n && (prefix+1)%10 != 0 {
			dfs(prefix + 1)
		}
	}
	dfs(1)
	return res
}

func lexicalOrder3(n int) (res []int) {
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

/*Самый быстрый*/
func lexicalOrder1(n int) []int {
	res := make([]int, n)
	index := 0
	i10 := 10
	i100 := 100
	i1000 := 1000
	i10000 := 10000
	for i := 1; i < 10 && i <= n; i++ {
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
					}
				}
			}
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

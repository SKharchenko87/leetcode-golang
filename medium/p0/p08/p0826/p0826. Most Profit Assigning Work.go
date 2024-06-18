package p0826

import (
	"slices"
	"sort"
)

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	n := len(difficulty)
	m := len(worker)
	job := make([][2]int, n)
	for i := 0; i < n; i++ {
		job[i][0] = difficulty[i]
		job[i][1] = profit[i]
	}
	slices.SortFunc(job, func(a, b [2]int) int { return a[0] - b[0] })
	sort.Ints(worker)
	resProfit := 0
	indexMax := 0
	maxProfit := 0
	for i := 0; i < m; i++ {
		for j := indexMax; j < n; j++ {
			if job[j][0] > worker[i] {
				break
			}
			if maxProfit < job[j][1] {
				maxProfit = job[j][1]
			}
			indexMax = j
		}
		resProfit += maxProfit
	}
	return resProfit
}

func maxProfitAssignment3(difficulty []int, profit []int, worker []int) int {
	n := len(difficulty)
	m := len(worker)
	job := make([][2]int, n)
	for i := 0; i < n; i++ {
		job[i][0] = difficulty[i]
		job[i][1] = profit[i]
	}
	sort.Slice(job, func(i, j int) bool {
		return job[i][0] < job[j][0]
	})
	sort.Ints(worker)
	resProfit := 0
	indexMax := 0
	maxProfit := 0
	for i := 0; i < m; i++ {
		for j := indexMax; j < n; j++ {
			if job[j][0] > worker[i] {
				break
			}
			if maxProfit < job[j][1] {
				maxProfit = job[j][1]
			}
			indexMax = j
		}
		resProfit += maxProfit
	}
	return resProfit
}

func maxProfitAssignment2(difficulty []int, profit []int, worker []int) int {
	n := len(difficulty)
	m := len(worker)

	var partition = func(low, high int) ([]int, int) {
		pivot := difficulty[high]
		i := low
		for j := low; j < high; j++ {
			if difficulty[j] < pivot {
				difficulty[i], difficulty[j] = difficulty[j], difficulty[i]
				profit[i], profit[j] = profit[j], profit[i]
				i++
			}
		}
		difficulty[i], difficulty[high] = difficulty[high], difficulty[i]
		profit[i], profit[high] = profit[high], profit[i]
		return difficulty, i
	}

	var quickSort func(low, high int) []int
	quickSort = func(low, high int) []int {
		if low < high {
			var p int
			difficulty, p = partition(low, high)
			difficulty = quickSort(low, p-1)
			difficulty = quickSort(p+1, high)
		}
		return difficulty
	}

	quickSort(0, len(difficulty)-1)

	sort.Ints(worker)

	resProfit := 0
	indexMax := 0
	maxProfit := 0
	for i := 0; i < m; i++ {
		for j := indexMax; j < n; j++ {
			if difficulty[j] > worker[i] {
				break
			}
			if maxProfit < profit[j] {
				maxProfit = profit[j]
			}
			indexMax = j
		}
		resProfit += maxProfit
	}
	return resProfit
}

func maxProfitAssignment1(difficulty []int, profit []int, worker []int) int {
	n := len(difficulty)
	m := len(worker)
	job := make(map[int]int, n)
	for i := 0; i < n; i++ {
		if v, ok := job[difficulty[i]]; ok && v < profit[i] || !ok {
			job[difficulty[i]] = profit[i]
		}
	}
	sort.Ints(difficulty)
	sort.Ints(worker)

	resProfit := 0
	indexMax := 0
	maxProfit := 0
	for i := 0; i < m; i++ {
		for j := indexMax; j < n; j++ {
			if difficulty[j] > worker[i] {
				break
			}
			if maxProfit < job[difficulty[j]] {
				maxProfit = job[difficulty[j]]
			}
			indexMax = j
		}
		resProfit += maxProfit
	}
	return resProfit
}

func maxProfitAssignment0(difficulty []int, profit []int, worker []int) int {
	n := len(difficulty)
	m := len(worker)
	job := make([][2]int, n)
	for i := 0; i < n; i++ {
		job[i][0] = difficulty[i]
		job[i][1] = profit[i]
	}
	sort.Slice(job, func(i, j int) bool {
		if job[i][0] > job[j][0] {
			return true
		} else if job[i][0] == job[j][0] {
			return job[i][1] > job[j][1]
		} else {
			return false
		}
	})

	sort.Slice(worker, func(i, j int) bool {
		return worker[i] > worker[j]
	})

	resProfit := 0
	indexMax := 0
	for i := 0; i < m; i++ {
		maxProfit := 0
		for j := indexMax; j < n; j++ {
			if job[j][0] <= worker[i] && maxProfit < job[j][1] {
				maxProfit = job[j][1]
				indexMax = j
			}
		}
		resProfit += maxProfit
	}
	return resProfit
}

//type job struct{ difficulty, profit int }
//
//func maxProfitAssignment(difficulty []int, profit []int, worker []int) (ans int) {
//	n := len(difficulty)
//	jobs := make([]job, n)
//	for i := range difficulty {
//		jobs[i] = job{difficulty: difficulty[i], profit: profit[i]}
//	}
//	// slices.SortFunc(jobs, func(a, b job) int { return a.difficulty - b.difficulty })
//	sort.Slice(jobs, func(i, j int) bool {
//		return jobs[i].difficulty < jobs[j].difficulty
//	})
//	slices.Sort(worker)
//	var mx, i int
//	for _, w := range worker {
//		for ; i < n && w >= jobs[i].difficulty; i++ {
//			mx = max(mx, jobs[i].profit)
//		}
//		ans += mx
//	}
//	return
//}

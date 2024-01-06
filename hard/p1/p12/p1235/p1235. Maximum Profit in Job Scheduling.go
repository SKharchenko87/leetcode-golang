package p1235

import "sort"

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	l := len(startTime)
	jobs := [][3]int{}
	for i := 0; i < l; i++ {
		jobs = append(jobs, [3]int{startTime[i], endTime[i], profit[i]})
	}

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][0] < jobs[j][0]
	})

	memo := map[int]int{}

	return dfs(0, memo, l, jobs)
}

func dfs(i int, memo map[int]int, length int, jobs [][3]int) int {
	if i == length {
		return 0
	}

	if v, ok := memo[i]; ok {
		return v
	}

	profit1 := dfs(i+1, memo, length, jobs)

	l, r := i+1, length
	for l < r {
		mid := l + (r-l)/2
		if jobs[mid][0] < jobs[i][1] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	profit2 := jobs[i][2] + dfs(l, memo, length, jobs)

	memo[i] = max(profit1, profit2)
	return memo[i]
}

package p3488

import "math"

func solveQueries(nums []int, queries []int) []int {
	nl := len(nums)
	ql := len(queries)
	preRes := make([]int, nl)
	lastIndexes := make(map[int]int, min(nl, ql))
	for i := 0; i < nl*3/2; i++ {
		currIndex := i % nl
		n := nums[currIndex]
		if preRes[currIndex] == 0 {
			preRes[currIndex] = math.MaxInt32
		}
		if lastIndex, ok := lastIndexes[n]; ok {
			lastIndex %= nl
			if cost := i - lastIndex; lastIndex >= 0 && cost > 0 {
				preRes[currIndex] = min(preRes[currIndex], cost)
				preRes[lastIndex] = min(preRes[lastIndex], cost)
			}
		}
		lastIndexes[n] = i
	}

	res := make([]int, ql)
	for i, query := range queries {
		if preRes[query] == math.MaxInt32 || preRes[query] == nl {
			res[i] = -1
		} else {
			res[i] = preRes[query]
		}

	}
	return res
}

func solveQueries2(nums []int, queries []int) []int {
	nl := len(nums)
	ql := len(queries)
	preRes := make([]int, nl)
	lastIndexes := make([]int, 1e6)
	for i := 0; i < nl*3/2; i++ {
		currIndex := i % nl
		n := nums[currIndex]
		lastIndex := (lastIndexes[n] - 1) % nl

		if preRes[currIndex] == 0 {
			preRes[currIndex] = math.MaxInt32
		}
		if cost := i - lastIndex; lastIndex >= 0 && cost > 0 {
			preRes[currIndex] = min(preRes[currIndex], cost)
			preRes[lastIndex] = min(preRes[lastIndex], cost)
		}

		lastIndexes[n] = i + 1
	}

	res := make([]int, ql)
	for i, query := range queries {
		if preRes[query] == math.MaxInt32 || preRes[query] == nl {
			res[i] = -1
		} else {
			res[i] = preRes[query]
		}

	}
	return res
}

func solveQueries1(nums []int, queries []int) []int {
	nl := len(nums)
	ql := len(queries)
	n := min(nl, ql)

	firstIndex := make(map[int]int, n)
	lastIndex := make(map[int]int, n)

	preAsk := make([]int, nl)
	for i := 0; i < nl; i++ {
		preAsk[i] = math.MaxInt32
		if first, ok := firstIndex[nums[i]]; ok {
			last := lastIndex[nums[i]]
			preAsk[i] = min(preAsk[i], i-last, nl-i+first)
			preAsk[last] = min(preAsk[last], i-last)
			preAsk[first] = min(preAsk[first], nl-i+first)
		} else {
			firstIndex[nums[i]] = i
		}
		lastIndex[nums[i]] = i
	}
	res := make([]int, ql)
	for i, query := range queries {
		if preAsk[query] == math.MaxInt32 {
			res[i] = -1
		} else {
			res[i] = preAsk[query]
		}

	}
	return res
}

func solveQueries0(nums []int, queries []int) []int {
	nl := len(nums)
	ql := len(queries)

	set := make(map[int]struct{}, min(nl, ql))
	for i := 0; i < ql; i++ {
		set[nums[queries[i]]] = struct{}{}
	}
	ul := len(set)

	preAsk := make([]int, nl)
	firstIndex := make(map[int]int, ul)
	for i := nl - 1; i >= 0; i-- {
		if _, ok := set[nums[i]]; ok {
			firstIndex[nums[i]] = i
			preAsk[i] = math.MaxInt32
		}
	}

	lastIndex := make(map[int]int, ul)

	for i := 0; i < nl; i++ {
		if _, ok := set[nums[i]]; ok {
			if v, exists := lastIndex[nums[i]]; exists {
				preAsk[v] = min(preAsk[v], i-v)
				preAsk[i] = min(preAsk[i], i-v)
			}
			lastIndex[nums[i]] = i
		}
	}

	for num, index := range lastIndex {
		if preAsk[index] == math.MaxInt32 {
			preAsk[index] = -1
		} else {
			fIndex := firstIndex[num]
			cost := nl - index + fIndex
			preAsk[index] = min(preAsk[index], cost)
			preAsk[fIndex] = min(preAsk[fIndex], cost)
		}
	}

	ask := make([]int, ql)
	for i := 0; i < ql; i++ {
		ask[i] = preAsk[queries[i]]
	}
	return ask
}

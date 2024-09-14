package p1310

func xorQueries(arr []int, queries [][]int) []int {
	lenArr, lenQueries := len(arr), len(queries)
	for i := 1; i < lenArr; i++ {
		arr[i] ^= arr[i-1]
	}

	result := make([]int, lenQueries)
	for i, query := range queries {
		if query[0] > 0 {
			result[i] = arr[query[0]-1] ^ arr[query[1]]
		} else {
			result[i] = arr[query[1]]
		}
	}
	return result
}

func xorQueries2(arr []int, queries [][]int) []int {
	lenArr, lenQueries := len(arr), len(queries)
	sums := make([]int, lenArr+1)
	sums[0] = 0
	for i := 0; i < lenArr; i++ {
		sums[i+1] = sums[i] ^ arr[i]
	}

	result := make([]int, lenQueries)
	for i, query := range queries {
		result[i] = sums[query[0]] ^ sums[query[1]+1]
	}
	return result
}

func xorQueries1(arr []int, queries [][]int) []int {
	m := make(map[int]int, len(queries))
	prefSum := 0
	for i := range arr {
		prefSum = prefSum ^ arr[i]
		m[i] = prefSum
	}
	result := make([]int, len(queries))
	for i, query := range queries {
		result[i] = m[query[0]-1] ^ m[query[1]]
	}
	return result
}

func xorQueries0(arr []int, queries [][]int) []int {
	result := make([]int, len(queries))
	for i, query := range queries {
		result[i] = arr[query[0]]
		for j := query[0] + 1; j <= query[1]; j++ {
			result[i] ^= arr[j]
		}
	}
	return result
}

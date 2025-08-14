package p3

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -1 * x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs64(x int64) int64 {
	if x >= 0 {
		return x
	}
	return -1 * x

}

func min64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func max64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func validSequence(word1 string, word2 string) []int {
	arr := [26][]int{}
	for i := 0; i < len(word1); i++ {
		if arr[int(word1[i]-'a')] == nil {
			arr[int(word1[i]-'a')] = []int{i}
		} else {
			arr[int(word1[i]-'a')] = append(arr[int(word1[i]-'a')], i)
		}
	}

	for i := 0; i < len(word2); i++ {
		if arr[int(word2[i]-'a')] == nil {
			return []int{}
		}
	}
	return []int{}
}

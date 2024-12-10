package p2554

func maxCount1(banned []int, n int, maxSum int) int {
	var sum, cnt int
	m := map[int]struct{}{}
	for _, b := range banned {
		if b <= n {
			m[b] = struct{}{}
		}
	}

	for i := 1; i <= n && sum <= maxSum; i++ {
		if _, ok := m[i]; !ok {
			sum += i
			cnt++
		}
	}
	if sum > maxSum {
		cnt--
	}
	return cnt
}

func maxCount(banned []int, n int, maxSum int) int {
	limit := min(n, maxSum)
	bannedMap := make([]bool, limit+1)
	for _, b := range banned {
		if b <= limit {
			bannedMap[b] = true
		}
	}
	var sum, cnt int
	for i := 1; i <= limit && sum <= maxSum; i++ {
		if !bannedMap[i] {
			sum += i
			cnt++
		}
	}
	if sum > maxSum {
		cnt--
	}
	return cnt
}

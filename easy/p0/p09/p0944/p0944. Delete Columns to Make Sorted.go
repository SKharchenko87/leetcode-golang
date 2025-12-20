package p0944

func minDeletionSize(strs []string) int {
	n := len(strs[0])
	cur := make([]byte, n)
	m := make([]bool, n)
	res := 0
	for _, str := range strs {
		for i := 0; i < n; i++ {
			if str[i] >= cur[i] {
				cur[i] = str[i]
			} else if !m[i] {
				m[i] = true
				res++
			}
		}
	}
	return res
}

func minDeletionSize1(strs []string) int {
	n := len(strs[0])
	cur := make([]bool, n)

	for j := 1; j < len(strs); j++ {
		for i := 0; i < n; i++ {
			if strs[j][i] < strs[j-1][i] {
				cur[i] = true
			}
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		if cur[i] {
			res++
		}
	}
	return res
}

func minDeletionSize0(strs []string) int {
	n := len(strs[0])
	cur := make([]byte, n)
	m := map[int]struct{}{}
	for _, str := range strs {
		for i := 0; i < n; i++ {
			if str[i] >= cur[i] {
				cur[i] = str[i]
			} else {
				m[i] = struct{}{}
			}
		}
	}
	return len(m)
}

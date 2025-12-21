package p0955

func minDeletionSize(strs []string) int {
	l := len(strs)
	n := len(strs[0])
	res := 0
	flags := make([]bool, n)
LOOP:
	for i := 1; i < l; i++ {
		for j := 0; j < n; j++ {
			if flags[j] || strs[i][j] == strs[i-1][j] {
				continue
			}
			if strs[i][j] < strs[i-1][j] {
				flags[j] = true
				res++
				i = 0
			}
			continue LOOP
		}
	}
	return res
}

func minDeletionSize0(strs []string) int {
	l := len(strs)
	n := len(strs[0])
	flags := make([]bool, n)
LOOP:
	for i := 1; i < l; i++ {
		for j := 0; j < n; j++ {
			if flags[j] || strs[i][j] == strs[i-1][j] {
				continue
			}
			if strs[i][j] < strs[i-1][j] {
				flags[j] = true
				i = 0
			}
			continue LOOP
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		if flags[i] {
			res++
		}
	}
	return res
}

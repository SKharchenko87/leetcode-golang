package p1415

func getHappyString(n int, k int) string {
	cur := make([]rune, n)
	var resStr string
	var dfs func(index int)
	dfs = func(index int) {
		if index == n {
			k--
			if k == 0 {
				resStr = string(cur)
			}
			return
		}
		for i := 'a'; i <= 'c'; i++ {
			if cur[index-1] == i {
				continue
			}
			cur[index] = i
			dfs(index + 1)
			if resStr != "" {
				return
			}
		}
	}
	for i := 'a'; i <= 'c'; i++ {
		cur[0] = i
		dfs(1)
	}
	return resStr
}

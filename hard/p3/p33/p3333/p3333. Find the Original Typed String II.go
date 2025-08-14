package p3333

const MOD int = 1e9 + 7

func possibleStringCount(word string, k int) int {
	n := len(word)
	cnt := 1
	var freq []int

	for i := 1; i < n; i++ {
		if word[i] == word[i-1] {
			cnt++
		} else {
			freq = append(freq, cnt)
			cnt = 1
		}
	}
	freq = append(freq, cnt)

	ans := 1
	for _, o := range freq {
		ans = ans * o % MOD
	}

	if len(freq) >= k {
		return ans
	}

	f := make([]int, k)
	g := make([]int, k)
	f[0] = 1
	for i := range g {
		g[i] = 1
	}

	for i := 0; i < len(freq); i++ {
		f_new := make([]int, k)
		for j := 1; j < k; j++ {
			f_new[j] = g[j-1]
			if j-freq[i]-1 >= 0 {
				f_new[j] = (f_new[j] - g[j-freq[i]-1] + MOD) % MOD
			}
		}
		g_new := make([]int, k)
		g_new[0] = f_new[0]
		for j := 1; j < k; j++ {
			g_new[j] = (g_new[j-1] + f_new[j]) % MOD
		}
		f, g = f_new, g_new
	}

	return (ans - g[k-1] + MOD) % MOD
}

// TLE
func possibleStringCount2(word string, k int) int {
	groups := []int{}
	cnt := 1
	for i := 1; i < len(word); i++ {
		if word[i] == word[i-1] {
			cnt++
		} else {
			groups = append(groups, cnt)
			cnt = 1
		}
	}
	groups = append(groups, cnt)

	counts := make([]int, len(word)+1)
	counts[0] = 1
	newCounts := make([]int, len(word)+1)
	for _, groupCnt := range groups {
		for i := 1; i <= len(word); i++ {
			if i-1 >= 0 {
				var x int
				if i-1-groupCnt >= 0 {
					x = counts[i-1-groupCnt]
				}
				newCounts[i] = (newCounts[i-1] + counts[i-1] - x + MOD) % MOD
			}
		}
		counts, newCounts = newCounts, counts
		clear(newCounts)
	}
	res := 0
	for l, cnt := range counts {
		if l >= k {
			res = (res + cnt) % MOD
		}
	}
	return res
}

// TLE
func possibleStringCount1(word string, k int) int {
	groups := []int{}
	cnt := 1
	for i := 1; i < len(word); i++ {
		if word[i] == word[i-1] {
			cnt++
		} else {
			groups = append(groups, cnt)
			cnt = 1
		}
	}
	groups = append(groups, cnt)

	counts := map[int]int{0: 1}
	newCounts := make(map[int]int)
	for _, groupCnt := range groups {
		for i := 1; i <= len(word); i++ {
			newCounts[i] = (newCounts[i-1] + counts[i-1] - counts[i-1-groupCnt] + MOD) % MOD
		}
		counts, newCounts = newCounts, counts
		clear(newCounts)
	}
	res := 0
	for l, cnt := range counts {
		if l >= k {
			res = (res + cnt) % MOD
		}
	}
	return res
}

// TLE
func possibleStringCount0(word string, k int) int {
	groups := []int{}
	cnt := 1
	for i := 1; i < len(word); i++ {
		if word[i] == word[i-1] {
			cnt++
		} else {
			groups = append(groups, cnt)
			cnt = 1
		}
	}
	groups = append(groups, cnt)

	counts := map[int]int{0: 1}
	newCounts := make(map[int]int)
	for _, groupCnt := range groups {
		for l, cnt := range counts {
			for i := 1; i <= groupCnt; i++ {
				newCounts[l+i] = (newCounts[l+i] + cnt) % MOD
			}
		}
		counts, newCounts = newCounts, counts
		clear(newCounts)
	}
	res := 0
	for l, cnt := range counts {
		if l >= k {
			res = (res + cnt) % MOD
		}
	}
	return res
}

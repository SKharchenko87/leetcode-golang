package p2275

func largestCombination(candidates []int) int {
	res := 0
	maxVal := 0
	c := 0
	for j := 0; j < len(candidates); j++ {
		if candidates[j]&1 != 0 {
			c++
		}
		if candidates[j] > maxVal {
			maxVal = candidates[j]
		}
	}
	if c > res {
		res = c
	}

	for i := 2; i < 1000_0000 && i <= maxVal; i *= 2 {
		c := 0
		for j := 0; j < len(candidates); j++ {
			if candidates[j]&i != 0 {
				c++
			}
		}
		if c > res {
			res = c
		}
	}
	return res
}

func largestCombination1(candidates []int) int {

	m := [24]int{}
	for j := 0; j < len(candidates); j++ {
		for i := range m {
			if candidates[j]&(1<<i) != 0 {
				m[i]++
			}
		}
	}

	res := 0
	for i := range m {
		if res < m[i] {
			res = m[i]
		}
	}

	return res
}

func largestCombination0(candidates []int) int {
	res := 0
	for i := 1; i < 1000_0000; i *= 2 {
		c := 0
		for j := 0; j < len(candidates); j++ {
			if candidates[j]&i != 0 {
				c++
			}
		}
		if c > res {
			res = c
		}
	}
	return res
}

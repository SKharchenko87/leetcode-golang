package p2125

func numberOfBeams(bank []string) int {
	var prev, cur, res int
	for _, row := range bank {
		cur = 0
		for _, c := range row {
			if c == '1' {
				cur++
			}
		}
		if cur > 0 {
			res += prev * cur
			prev = cur
		}
	}
	return res
}

func numberOfBeams0(bank []string) int {
	levels := len(bank)
	res := 0
	prevCount := 0
	for i := 0; i < levels; i++ {
		count := 0
		for j := 0; j < len(bank[i]); j++ {
			if bank[i][j] == 49 {
				count++
			}
		}
		if prevCount == 0 {
			prevCount = count
		} else if count != 0 {
			res = res + prevCount*count
			prevCount = count
		}
	}
	return res
}

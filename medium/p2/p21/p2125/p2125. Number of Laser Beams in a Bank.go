package p2125

func numberOfBeams(bank []string) int {
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

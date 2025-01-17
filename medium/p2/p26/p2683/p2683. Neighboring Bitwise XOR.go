package p2683

func doesValidArrayExist(derived []int) bool {
	res := 0
	for _, v := range derived {
		res ^= v
	}
	return res == 0
}

func doesValidArrayExist1(derived []int) bool {
	cnt := 0
	for _, v := range derived {
		cnt += v
	}
	return cnt&1 == 0
}

func doesValidArrayExist0(derived []int) bool {
	res := true
	for _, v := range derived {
		if v == 1 {
			res = !res
		}
	}
	return res
}

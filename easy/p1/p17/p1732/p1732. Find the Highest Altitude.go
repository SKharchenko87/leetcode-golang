package p1732

func largestAltitude(gain []int) int {
	res, cur := 0, 0
	for i := 0; i < len(gain); i++ {
		cur += gain[i]
		res = max(res, cur)
	}
	return res
}

func largestAltitude0(gain []int) int {
	maxAltitude, certitude := 0, 0
	for _, v := range gain {
		certitude += v
		if certitude > maxAltitude {
			maxAltitude = certitude
		}
	}
	return maxAltitude
}

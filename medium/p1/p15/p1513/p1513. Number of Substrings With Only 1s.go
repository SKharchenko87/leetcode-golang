package p1513

const MOD = 1e9 + 7

func numSub(s string) int {
	sub := 0
	res := 0
	for _, ch := range s {
		if ch == '0' {
			res += sub * (sub + 1) / 2
			res %= MOD
			sub = 0
		} else {
			sub++
		}
	}
	res += sub * (sub + 1) / 2
	res %= MOD
	return res
}

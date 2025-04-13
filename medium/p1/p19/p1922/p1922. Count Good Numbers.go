package p1922

const MOD = 1e9 + 7

func countGoodNumbers(n int64) int {
	return int(powQuick(5, (n+1)/2) * powQuick(4, n/2) % MOD)
}

func powQuick(x, y int64) int64 {
	res := int64(1)
	for y != 0 {
		if y%2 == 1 {
			res = res * x % MOD
		}
		x = x * x % MOD
		y /= 2
	}
	return res
}

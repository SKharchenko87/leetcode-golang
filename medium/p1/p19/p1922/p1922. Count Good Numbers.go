package p1922

const MOD = 1e9 + 7

func countGoodNumbers(n int64) int {
	return int(powerWithMOD64(5, (n+1)/2, MOD) * powerWithMOD64(4, n/2, MOD) % MOD)
}

func powerWithMOD64(x, y, mod int64) int64 {
	res := int64(1)
	for y != 0 {
		if y%2 == 1 {
			res = res * x % mod
		}
		x = x * x % mod
		y /= 2
	}
	return res
}

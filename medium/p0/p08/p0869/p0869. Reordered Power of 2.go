package p0869

var digitsPowerOf2 = map[[10]int]struct{}{}

func init() {
	for i := 0; i < 30; i++ {
		tmp := 1 << i
		digitsPowerOf2[countDigits(tmp)] = struct{}{}
	}
}

func countDigits(n int) (result [10]int) {
	for n > 0 {
		result[n%10]++
		n /= 10
	}
	return result
}

func reorderedPowerOf2(n int) bool {
	_, exists := digitsPowerOf2[countDigits(n)]
	return exists
}

package p3345

func smallestNumber(n int, t int) int {
	for ; n%10 > 0; n++ {
		if productOfDigits(n)%t == 0 {
			break
		}
	}
	return n
}

func productOfDigits(n int) int {
	res := 1
	for n > 0 && res > 0 {
		res *= n % 10
		n /= 10
	}
	return res
}

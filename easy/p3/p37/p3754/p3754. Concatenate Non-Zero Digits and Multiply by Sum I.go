package p3754

func sumAndMultiply(n int) int64 {
	d, s := solution(n)
	return int64(d * s)
}

func solution(n int) (res, sum int) {
	for n > 0 {
		if x := n % 10; x > 0 {
			sum += x
			defer func(x int) {
				res = res*10 + x%10
			}(x)
		}
		n /= 10
	}
	return
}

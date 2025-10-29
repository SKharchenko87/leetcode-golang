package p3370

func smallestNumber(n int) (res int) {
	for n > res {
		res = res<<1 + 1
	}
	return res
}

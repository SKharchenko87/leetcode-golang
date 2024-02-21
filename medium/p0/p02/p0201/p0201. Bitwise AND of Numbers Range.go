package p0201

func rangeBitwiseAnd(left int, right int) int {
	cnt := 0
	for left != right {
		cnt++
		left >>= 1
		right >>= 1
	}
	return left << cnt
}

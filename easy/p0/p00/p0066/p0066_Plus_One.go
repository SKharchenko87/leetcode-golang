package p0066

func plusOne(digits []int) []int {
	l := len(digits)
	p := 1
	for i := l - 1; i >= 0 && p == 1; i-- {
		digits[i] += p
		p = digits[i] / 10
		digits[i] %= 10
	}
	if p == 1 {
		return append([]int{1}, digits...)
	}
	return digits
}

func plusOne0(digits []int) []int {
	one := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i], one = (digits[i]+one)%10, (digits[i]+one)/10
	}
	if one == 1 {
		res := []int{1}
		return append(res, digits...)
	}
	return digits
}

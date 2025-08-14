package p0650

func minSteps(n int) int {
	if n == 1 {
		return 0
	}
	res := 0
	i := 2
	for n != 1 {
		for n%i != 0 {
			i++
		}
		res += i
		n /= i
	}
	return res
}

func numberDivisors(n int) []int {
	res := make([]int, 0)
	i := 2
	for n != 1 {
		for n%i != 0 {
			i++
		}
		res = append(res, i)
		n /= i
	}
	return res
}

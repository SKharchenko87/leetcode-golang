package p3751

func totalWaviness(num1 int, num2 int) int {
	res := 0
	for i := num1; i <= num2; i++ {
		res += countWaviness(i)
	}
	return res
}

func countWaviness(n int) int {
	if n < 101 {
		return 0
	}
	prev := n % 10
	n /= 10
	curr := n % 10
	n /= 10
	cnt := 0
	for n > 0 {
		next := n % 10
		if prev < curr && curr > next || prev > curr && curr < next {
			cnt++
		}
		prev, curr = curr, next
		n /= 10
	}
	return cnt
}

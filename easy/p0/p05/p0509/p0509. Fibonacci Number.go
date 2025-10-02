package p0509

const MAX_N = 30

var fibonacci = make([]int, MAX_N+1)

func init() {
	fibonacci[0] = 0
	fibonacci[1] = 1
	for i := 2; i <= MAX_N; i++ {
		fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
	}
}

func fib(n int) int {
	return fibonacci[n]
}

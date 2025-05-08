package p0790

const mod = 1e9 + 7

const size = 1000

var result = [size]int{}

func numTilings(n int) int {
	return result[n-1]
}

func init() {
	result[0] = 1
	result[1] = 2
	result[2] = 5
	for i := 3; i < size; i++ {
		result[i] = (result[i-1]*2 + result[i-3]) % mod
	}
}

func numTilings0(n int) int {
	var l int
	if n < 4 {
		l = 4
	} else {
		l = n
	}
	arr := make([]int, l)
	arr[0] = 1
	arr[1] = 2
	arr[2] = 5
	for i := 3; i < n; i++ {
		arr[i] = (arr[i-1]*2 + arr[i-3]) % mod
	}
	return arr[n-1]
}

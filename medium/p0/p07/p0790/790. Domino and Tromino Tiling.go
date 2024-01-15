package p0790

const mod = 1e9 + 7

func numTilings(n int) int {
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

package p0231

import (
	"math/bits"
	"slices"
)

/*
Benchmark_isPowerOfTwo-12     	1000000000	         0.2871 ns/op	       0 B/op	       0 allocs/op
Benchmark_isPowerOfTwo5-12    	79230408	        16.76 ns/op	       0 B/op	       0 allocs/op
Benchmark_isPowerOfTwo4-12    	87269553	        14.03 ns/op	       0 B/op	       0 allocs/op
Benchmark_isPowerOfTwo3-12    	1000000000	         0.5969 ns/op	       0 B/op	       0 allocs/op
Benchmark_isPowerOfTwo2-12    	1000000000	         0.2775 ns/op	       0 B/op	       0 allocs/op
Benchmark_isPowerOfTwo1-12    	59751135	        17.79 ns/op	       0 B/op	       0 allocs/op
*/

var powerOfTwo = []int{1, 1 << 1, 1 << 2, 1 << 3, 1 << 4, 1 << 5, 1 << 6, 1 << 7, 1 << 8, 1 << 9, 1 << 10, 1 << 11, 1 << 12, 1 << 13, 1 << 14, 1 << 15, 1 << 16, 1 << 17, 1 << 18, 1 << 19, 1 << 20, 1 << 21, 1 << 22, 1 << 23, 1 << 24, 1 << 25, 1 << 26, 1 << 27, 1 << 28, 1 << 29, 1 << 30, 1 << 31}
var powerOfTwoSet = map[int]struct{}{1: {}, 1 << 1: {}, 1 << 2: {}, 1 << 3: {}, 1 << 4: {}, 1 << 5: {}, 1 << 6: {}, 1 << 7: {}, 1 << 8: {}, 1 << 9: {}, 1 << 10: {}, 1 << 11: {}, 1 << 12: {}, 1 << 13: {}, 1 << 14: {}, 1 << 15: {}, 1 << 16: {}, 1 << 17: {}, 1 << 18: {}, 1 << 19: {}, 1 << 20: {}, 1 << 21: {}, 1 << 22: {}, 1 << 23: {}, 1 << 24: {}, 1 << 25: {}, 1 << 26: {}, 1 << 27: {}, 1 << 28: {}, 1 << 29: {}, 1 << 30: {}, 1 << 31: {}}

func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}

func isPowerOfTwo5(n int) bool {
	_, ok := powerOfTwoSet[n]
	return ok
}

func isPowerOfTwo4(n int) bool {
	_, exists := slices.BinarySearch(powerOfTwo, n)
	return exists
}

func isPowerOfTwo3(n int) bool {
	return bits.OnesCount(uint(n)) == 1
}

func isPowerOfTwo2(n int) bool {
	switch n {
	case 1, 1 << 1, 1 << 2, 1 << 3, 1 << 4, 1 << 5, 1 << 6, 1 << 7, 1 << 8, 1 << 9, 1 << 10, 1 << 11, 1 << 12, 1 << 13, 1 << 14, 1 << 15, 1 << 16, 1 << 17, 1 << 18, 1 << 19, 1 << 20, 1 << 21, 1 << 22, 1 << 23, 1 << 24, 1 << 25, 1 << 26, 1 << 27, 1 << 28, 1 << 29, 1 << 30, 1 << 31:
		return true
	}
	return false
}

func isPowerOfTwo2(n int) bool {
	if n <= 0 {
		return false
	}
	x := 1
	for n > x {
		x = x << 1
	}
	return n == x
}

func isPowerOfTwo1(n int) bool {
	s := 0
	if n < 0 {
		return false
	}
	for n != 0 {
		if n%2 == 1 {
			s++
			if s > 1 {
				return false
			}
		}
		n = n >> 1
	}
	return s == 1
}

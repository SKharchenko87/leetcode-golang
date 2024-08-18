package p0264

import (
	"fmt"
)

func nthUglyNumber(n int) int {
	if n <= 6 {
		return n
	}
	uglyNumber := make([]int, n)
	uglyNumber[0] = 1
	index2, index3, index5 := 0, 0, 0
	n2, n3, n5 := 2, 3, 5
	for i := 1; i < n; i++ {
		cur := min(n2, n3, n5)
		uglyNumber[i] = cur
		if cur == n2 {
			index2++
			n2 = uglyNumber[index2] * 2
		}
		if cur == n3 {
			index3++
			n3 = uglyNumber[index3] * 3
		}
		if cur == n5 {
			index5++
			n5 = uglyNumber[index5] * 5
		}
	}
	return uglyNumber[n-1]
}

func nthUglyNumber1(n int) int {
	if n <= 6 {
		return n
	} else if n < 10 {
		return n + 1
	}
	n -= 10
	x := n / 22
	y := n % 22
	m := map[int]int{0: 2, 1: 4, 2: 5, 3: 6, 4: 8, 5: 10, 6: 11, 7: 12, 8: 14, 9: 15, 10: 16,
		11: 17, 12: 18, 13: 20, 14: 22, 15: 23, 16: 24, 17: 25, 18: 26, 19: 28, 20: 29, 21: 30}
	return x*30 + m[y] + 10
}

func nthUglyNumber0(n int) int {
	i := 1
	n--
	for n != 0 {
		i++
		if i%2 == 0 || i%3 == 0 || i%5 == 0 {
			n--
		}
	}
	return i
}

func run() {
	i := 1
	n := 1
	fmt.Println("n=", n, " i=", i, " x=", nthUglyNumber(n))
	for n < 1600 {
		i++
		if i%2 == 0 || i%3 == 0 || i%5 == 0 {
			n++
			x := nthUglyNumber(n)
			if x != i {
				fmt.Println("n=", n, " i=", i, " x=", x)
			}
		}
	}
}

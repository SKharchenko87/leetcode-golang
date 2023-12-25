package main

import "fmt"

func mySqrt(x int) int {
	if x == 0 {
		return x
	}
	if x < 4 {
		return 1
	}
	if x < 9 {
		return 2
	}
	start := 3
	end := x / 3
	m := start + (end-start)/2
	for start <= end {
		if m*m == x {
			return m
		} else if m*m < x {
			start = m + 1
		} else {
			end = m - 1
		}
		m = start + (end-start)/2
	}
	return m - 1
}

func main() {
	fmt.Println(mySqrt(120))
}

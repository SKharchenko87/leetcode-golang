package main

import "fmt"

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	m := map[int]int{1: 1, 2: 2}
	for i := 3; i <= n; i++ {
		m[i] = m[i-1] + m[i-2]
	}
	return m[n]
}

func main() {
	fmt.Println(climbStairs(5))
}

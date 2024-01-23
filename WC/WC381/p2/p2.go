package p2

import "fmt"

func countOfPairs(n int, x int, y int) []int {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[n-i-1] = i * 2
	}
	fmt.Println(res)
	//arr := make([]int, n)

	return res
}

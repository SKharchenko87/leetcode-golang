package p1823

func findTheWinner(n int, k int) int {
	res := 0
	for i := 2; i <= n; i++ {
		res = (res + k) % i
	}
	return res + 1
}

func findTheWinner0(n int, k int) int {
	win := make([]int, n)
	for i := 0; i < n; i++ {
		win[i] = i
	}
	index := 0
	for i := n; i > 1; i-- {
		optimizeK := (k - 1) % n
		index = (index + optimizeK) % len(win)
		win = append(win[:index], win[index+1:]...)
	}
	return win[0] + 1
}

// 6 5
// 1 2 3 4 5 6
// 1 2 3 4 6
// 1 2 3 6
// 1 2 3
// 1 3
// 1

// 6 5
// 1 2 3 4 5 6 - 5%2 = 1 res=1
// 1 2 3 4 6 - 6%3 = 0 res=1
// 1 2 3 6 - 6%4 = 2 res=3
// 1 2 3 - 8%5 = 3 res=1
// 1 3  - 6%6 = 0 res=0

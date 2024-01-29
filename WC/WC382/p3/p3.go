package p3

import "fmt"

func flowerGame(n int, m int) int64 {
	s := n + m
	var count int64 = 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if (s-i-j)%2 == 0 && !(n%2 == 0 && m%2 == 0 && i == 1 && j == 1) {
				fmt.Println(i, j)
				count++
			}
		}
	}
	return count
}

func flowerGameEdu(n int, m int) int64 {
	var count int64 = 0
	for i := 1; i <= n; i++ {
		if i&1 == 1 {
			count += int64(m >> 1)
		} else {
			count += int64((m + 1) >> 1)
		}
	}
	return count
}

/*
class Solution:
    def flowerGame(self, n: int, m: int) -> int:
        ans = 0
        for x in range(1, n + 1):
            if x & 1:
                ans += m >> 1
            else:
                ans += (m + 1) >> 1
        return ans
*/

package p1458

import (
	"slices"
)

/*

     2   1  -2   5
 3   6   3  -6  15
 0   0   0   0   0
-6 -12  -6  12 -30


 3   6   3  -6  15
 0   6   6   6  15
-6   6   6  18  15


     2  -6   7
 3   6 -18  21
-2 -12  12 -14


*/

// ToDo можно улучшить - меняем местами n и m, и проходим построчно держа в памяти только предыдущую и текущую строку
func maxDotProduct(nums1 []int, nums2 []int) int {
	n, m := len(nums1), len(nums2)
	if n > m {
		nums1, nums2 = nums2, nums1
		n, m = m, n
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	dp[0][0] = nums1[0] * nums2[0]
	for j := 1; j < m; j++ {
		dp[0][j] = max(nums1[0]*nums2[j], dp[0][j-1])
	}
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], nums1[i]*nums2[0])
		for j := 1; j < m; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]+nums1[i]*nums2[j], nums1[i]*nums2[j])
		}
	}
	return slices.Max(dp[n-1])
}

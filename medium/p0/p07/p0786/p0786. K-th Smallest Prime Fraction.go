package p0786

import "sort"

func kthSmallestPrimeFraction(arr []int, k int) []int {
	n := len(arr)
	bigArr := make([][2]int, (n+1)*n/2-n)
	bigIndex := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			bigArr[bigIndex] = [2]int{arr[i], arr[j]}
			bigIndex++
		}
	}
	sort.Slice(bigArr, func(i, j int) bool {
		return bigArr[i][0]*bigArr[j][1] < bigArr[i][1]*bigArr[j][0]
	})
	return []int{bigArr[k-1][0], bigArr[k-1][1]}
}

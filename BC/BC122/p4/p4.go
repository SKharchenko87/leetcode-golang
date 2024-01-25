package p4

import (
	"fmt"
	"sort"
)

func minimumCost(nums []int, k int, dist int) int64 {
	l := len(nums)
	arr := make([][2]int, l)
	for i, v := range nums {
		arr[i] = [2]int{i, v}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] < arr[j][1]
	})
	fmt.Println(arr)

	return 0
}

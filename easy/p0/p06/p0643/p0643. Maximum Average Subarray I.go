package main

import (
	"fmt"
)

func findMaxAverage(nums []int, k int) float64 {

	sum := 0
	i := 0
	for ; i < k; i++ {
		sum += nums[i]
	}
	max := sum
	for i := k; i < len(nums); i++ {
		sum -= nums[i-k]
		sum += nums[i]
		if max < sum {
			max = sum
		}
	}
	return float64(max) / float64(k)
}

func main() {
	nums, k := []int{1, 12, -5, -6, 50, 3}, 4
	fmt.Println(findMaxAverage(nums, k))
}

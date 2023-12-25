package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{5, 4, 2, 3}
	fmt.Println(numberGame(nums))
}

func numberGame(nums []int) []int {
	res := make([]int, len(nums))
	sort.Ints(nums)
	for i := 0; i < len(nums); i = i + 2 {
		res[i], res[i+1] = nums[i+1], nums[i]
	}
	return res
}

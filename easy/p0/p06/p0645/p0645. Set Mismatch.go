package main

import (
	"fmt"
)

func findErrorNums(nums []int) []int {
	m := make(map[int]int, len(nums))
	ind := -1
	cur := 0
	curCount := 1
	for i := 1; i < len(nums); i++ {
		if cur == nums[i] {

		}
	}
	return int{0, 0}
}

func main() {
	nums := []int{1, 2, 2, 4}
	fmt.Println(findErrorNums(nums))
}

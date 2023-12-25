package main

import (
	"fmt"
	"sort"
)

func divideArray(nums []int, k int) [][]int {
	sort.Ints(nums)
	res := make([][]int, len(nums)/3)
	for i := 0; i < len(nums); i = i + 3 {
		if nums[i+1]-nums[i] > k || nums[i+2]-nums[i] > k {
			return [][]int{}
		}
		res[i/3] = []int{nums[i], nums[i+1], nums[i+2]}
	}
	return res
}

func main() {
	//nums := []int{1, 3, 4, 8, 7, 9, 3, 5, 1}
	//nums := []int{1, 3, 3, 2, 7, 3}
	nums := []int{1, 3, 4, 8, 7, 9, 3, 5, 1}
	fmt.Println(nums)
	//fmt.Println(divideArray(nums, 2))
	//fmt.Println(divideArray(nums, 3))
	fmt.Println(divideArray(nums, 2))
	//], k = 2
}

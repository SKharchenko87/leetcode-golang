package main

import (
	"fmt"
	"sort"
)

func main() {
	//nums := []int{1, 12, 1, 2, 5, 50, 3}
	nums := []int{1, 1, 2, 3, 5, 12, 50}
	//nums := []int{5, 5, 1, 50}
	//nums := []int{5, 5, 50}
	//nums := []int{5, 5, 5}
	//nums := []int{1, 1, 1, 50, 51, 100}
	//nums := []int{5, 5, 5, 50}
	fmt.Println(largestPerimeter(nums))

}

func largestPerimeter(nums []int) int64 {
	var sum int64
	sort.Ints(nums)
	for j := len(nums) - 1; j >= 0; j-- {
		sum = sum + int64(nums[j])
	}

	i := len(nums) - 1
	for ; i >= 2; i-- {
		v := nums[i]
		sum = sum - int64(v)
		if sum > int64(v) {
			sum = sum + int64(v)
			break
		}
	}
	if i <= 1 {
		return -1
	}
	return sum
}

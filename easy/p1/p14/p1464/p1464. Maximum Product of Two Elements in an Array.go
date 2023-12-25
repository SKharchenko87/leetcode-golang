package main

import "fmt"

func maxProduct(nums []int) int {
	i := 0
	j := 1
	for k := 2; k < len(nums); k++ {
		if nums[k] > nums[i] {
			if nums[i] > nums[j] {
				j = i
			}
			i = k
		}
		if i != k && nums[k] > nums[j] {
			j = k
		}
	}
	return (nums[i] - 1) * (nums[j] - 1)
}

func main() {
	nums := []int{2, 2, 1, 8, 1, 5, 4, 5, 2, 10, 3, 6, 5, 2, 3}
	fmt.Println(maxProduct(nums))
}

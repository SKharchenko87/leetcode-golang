package main

import "fmt"

func productExceptSelf(nums []int) []int {
	l := len(nums)
	res := make([]int, l)
	left := 1
	right := 1
	for i := 0; i < l; i++ {
		res[i] = 1
	}
	for i := 0; i < l; i++ {
		res[i], left = res[i]*left, left*nums[i]

		j := l - 1 - i
		res[j], right = res[j]*right, right*nums[j]
	}
	return res
}

func main() {
	//nums := []int{1, 2, 3, 4}
	nums := []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(nums))
}

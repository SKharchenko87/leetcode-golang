package main

import "fmt"

func pivotIndex(nums []int) int {
	fullSum := 0
	for _, v := range nums {
		fullSum += v
	}
	leftSum := 0
	rightSum := fullSum
	for i, v := range nums {
		rightSum = fullSum - leftSum
		leftSum += v
		if leftSum == rightSum {
			return i
		}
	}
	return -1
}

func main() {
	//nums := []int{1, 7, 3, 6, 5, 6}
	//nums := []int{1, 2, 3}
	nums := []int{2, 1, -1}
	fmt.Println(pivotIndex(nums))
}

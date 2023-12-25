package main

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
	x, y := math.MaxInt, math.MaxInt
	for _, z := range nums {
		if z <= x {
			x = z
		} else if z <= y {
			x = y
		} else {
			return true
		}
	}
	return false
}

func main() {
	//nums := []int{1, 2, 3, 4, 5}
	//nums := []int{20, 100, 10, 12, 5, 13}
	//nums := []int{13, 20, 100, 10, 12, 5, 13}
	//nums := []int{2, 1, 5, 0, 4, 6}
	nums := []int{9, 10, 5, 11, 10, 9, 8}
	//nums := []int{1,1,1,1,1,1,3,7}
	//nums := []int{9, 10, 5, 11, 10, 9, 8}
	fmt.Println(increasingTriplet(nums))
}

package main

import (
	"fmt"
	"math"
)

func maxProductDifference(nums []int) int {
	w, x, y, z := math.MinInt, math.MinInt, math.MaxInt, math.MaxInt
	for _, v := range nums {
		if v >= w {
			w, x = v, w
		} else if v >= x {
			x = v
		}
		if v <= y {
			y, z = v, y
		} else if v <= z {
			z = v
		}
	}
	fmt.Println(w, x, y, z)
	return w*x - y*z
}

func main() {
	nums := []int{5, 6, 2, 7, 4}
	fmt.Println(maxProductDifference(nums))
}

package main

import "fmt"

func maxArea(height []int) int {
	max := 0
	l := 0
	r := len(height) - 1
	for l < r {
		x := (r - l) * min(height[l], height[r])
		if x > max {
			max = x
		}
		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}
	return max
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}

package main

import (
	"fmt"
	"math"
)

func findSpecialInteger(arr []int) int {
	cur := arr[0]
	curCount := 1
	x := arr[0]
	xCount := 1
	for i := 1; i < len(arr); i++ {
		if cur == arr[i] {
			curCount++
		} else {
			if curCount >= xCount {
				xCount = curCount
				x = cur
			}
			cur = arr[i]
			curCount = 1
		}
	}
	if curCount >= xCount {
		xCount = curCount
		x = cur
	}
	return x
}

func main() {
	arr := []int{1, 2, 3, 3}
	fmt.Println(math.Ceil(float64(len(arr)) / 4))
	fmt.Println(findSpecialInteger(arr))
}

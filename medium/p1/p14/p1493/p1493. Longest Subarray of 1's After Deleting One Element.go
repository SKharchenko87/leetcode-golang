package main

import "fmt"

func longestSubarray(nums []int) int {
	prevLength := 0
	cursorZero := -1
	longest := 0
	for i, v := range nums {
		if v == 0 {
			length := i - cursorZero - 1
			if longest < prevLength+length {
				longest = prevLength + length
			}
			prevLength = length
			cursorZero = i
		}
	}
	length := len(nums) - cursorZero - 1
	if longest < prevLength+length {
		longest = prevLength + length
	}
	if cursorZero == -1 {
		longest = longest - 1
	}
	return longest
}

func main() {
	//nums := []int{0, 1, 1, 1, 0, 1, 1, 0, 1}
	nums := []int{1, 1, 1}
	//nums := []int{1, 1, 0, 1}
	//nums := []int{1, 1, 0, 0, 1, 1, 1, 0, 1}
	fmt.Println(longestSubarray(nums))
}

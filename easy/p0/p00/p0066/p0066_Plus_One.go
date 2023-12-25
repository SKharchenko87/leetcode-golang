package main

import "fmt"

func plusOne(digits []int) []int {
	//res := make([]int, len(digits)+1)
	one := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i], one = (digits[i]+one)%10, (digits[i]+one)/10
	}
	if one == 1 {
		res := []int{1}
		return append(res, digits...)
	}
	return digits
}

func main() {
	//digits := []int{4, 3, 2, 1}
	digits := []int{9, 9, 9, 9}
	fmt.Println(plusOne(digits))
}

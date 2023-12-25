package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	t := false
	if x < 0 {
		x = -1 * x
		t = true
	}
	res := 0

	for x != 0 {
		res = res*10 + x%10
		x = x / 10
	}
	if t {
		res = -1 * res
	}
	if res > int(math.Pow(float64(-2), float64(31))) && res < int(math.Pow(float64(2), float64(31)))-1 {
		return res
	} else {
		return 0
	}
	return res
}

func main() {
	x := -2147483412
	fmt.Println(reverse(x))
}

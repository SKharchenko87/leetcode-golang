package main

import (
	"fmt"
	"slices"
)

func kidsWithCandies(candies []int, extraCandies int) []bool {
	res := make([]bool, len(candies))
	maxCandie := slices.Max(candies)
	for i, v := range candies {
		res[i] = v+extraCandies >= maxCandie
	}
	return res
}

func main() {
	candies, extraCandies := []int{2, 3, 5, 1, 3}, 3
	fmt.Println(kidsWithCandies(candies, extraCandies))
}

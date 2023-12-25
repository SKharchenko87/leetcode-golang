package main

import (
	"fmt"
)

func buyChoco(prices []int, money int) int {
	min0 := 101
	min1 := 101
	for _, v := range prices {
		if min0 > v {
			min0, min1 = v, min0
		} else if min1 > v {
			min1 = v
		}
	}
	min := min0 + min1
	if min <= money {
		return money - min
	}
	return money
}

func main() {
	//prices, money := []int{3, 2, 3}, 3
	prices, money := []int{69, 91, 78, 19, 40, 13}, 94
	fmt.Println(buyChoco(prices, money))
}

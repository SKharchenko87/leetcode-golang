package main

import "fmt"

func largestAltitude(gain []int) int {
	maxAltitude, curltitude := 0, 0
	for _, v := range gain {
		curltitude += v
		if curltitude > maxAltitude {
			maxAltitude = curltitude
		}
	}
	return maxAltitude
}

func main() {
	//gain := []int{-5, 1, 5, 0, -7}
	gain := []int{-4, -3, -2, -1, 4, 3, 2}
	fmt.Println(largestAltitude(gain))
}

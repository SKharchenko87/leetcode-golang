package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	k := 0
	buf := 0
	flowerbed = append([]int{0}, flowerbed...)
	flowerbed = append(flowerbed, 0)
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 1 {
			if buf > 2 {
				fmt.Println(i, k, buf)
				k += (buf - 1) / 2
				if k >= n {
					return true
				}
			}
			buf = 0
		} else {
			buf++
		}
	}
	k += (buf - 1) / 2
	return k >= n
}

func main() {
	//flowerbed, n := []int{1, 0, 0, 0, 1, 0, 0}, 2
	flowerbed, n := []int{0, 1, 0}, 1
	//flowerbed, n := []int{1, 0, 0, 0, 0, 1}, 2
	fmt.Println(canPlaceFlowers(flowerbed, n))
}

package main

import "fmt"

func hammingWeight(num uint32) int {
	var b int = 0
	for num != 0 {
		if num%2 == 1 {
			b++
		}
		num = num >> 1
	}
	return b
}

func main() {
	var x uint32 = 0b00000000000000000000000000001011
	fmt.Println(hammingWeight(x))
}

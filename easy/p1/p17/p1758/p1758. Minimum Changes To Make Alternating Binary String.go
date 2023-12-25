package main

import "fmt"

func main() {
	s := "0100"
	fmt.Println(minOperations(s))

}

func minOperationsC(s string, c int32) int {
	cnt := 0
	for _, v := range s {
		if v-'0' != c {
			cnt++
		}
		c = (c + 1) % 2
	}
	return cnt
}

func minOperations1(s string) int {
	return 0
}

func minOperations(s string) int {
	var x int32
	c0, c1 := int32(0), int32(1)
	cnt0, cnt1 := 0, 0
	for _, v := range s {
		x = v - '0'
		if x != c0 {
			cnt0++
		}
		c0 ^= 1
		if x != c1 {
			cnt1++
		}
		c1 ^= 1
	}
	return min(cnt0, cnt1)
}

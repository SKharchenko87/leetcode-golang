package p2154

import (
	"math/bits"
	"sort"
)

func findFinalValue(nums []int, original int) int {
	var p, pow2 uint
	for _, num := range nums {
		pow2 = uint(num / original)
		if num%original == 0 && bits.OnesCount(pow2) == 1 {
			p |= pow2
		}
	}
	p++
	return original * 1 << bits.TrailingZeros(p)
}

var p = make([]bool, 11)

func findFinalValue1(nums []int, original int) int {
	clear(p)

	var pow2 uint
	for _, num := range nums {
		pow2 = uint(num / original)
		if num%original == 0 && bits.OnesCount(pow2) == 1 {
			p[bits.TrailingZeros(pow2)] = true
		}
	}
	for i, b := range p {
		if !b {
			return original * 1 << i
		}
	}
	return -1
}

func findFinalValue0(nums []int, original int) int {
	sort.Ints(nums)
	for _, num := range nums {
		if num == original {
			original *= 2
		} else if num > original {
			return original
		}
	}
	return original
}

func findFinalValueX(nums []int, original int) int {
	bits := 0
	for _, num := range nums {
		if num%original != 0 {
			continue
		}
		n := num / original
		if n&(n-1) == 0 {
			bits |= n
		}
	}
	bits++
	return original * (bits & -bits)
}

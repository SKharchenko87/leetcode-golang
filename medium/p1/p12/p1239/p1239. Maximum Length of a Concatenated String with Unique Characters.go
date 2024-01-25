package p1239

import (
	"math/bits"
)

func maxLength(arr []string) int {
	l := len(arr)
	arrBit := make([]uint32, l)
	for i, str := range arr {
		tmp := strToBit(str)
		if tmp != 0 {
			arrBit[i] = tmp
		}
	}
	maxBit := 0
	t := make(map[uint32][]uint32, l)
	for i := 0; i < l; i++ {
		curStr := arrBit[i]
		t[curStr] = []uint32{curStr}
		maxBit = max(maxBit, bits.OnesCount32(curStr))
		for j := 0; j < i; j++ {
			for _, x := range t[arrBit[j]] {
				if x&curStr == 0 {
					t[curStr] = append(t[curStr], x+curStr)
					maxBit = max(maxBit, bits.OnesCount32(x+curStr))
				}
			}
		}
	}
	return maxBit
}

func strToBit(str string) uint32 {
	var res uint32
	m := make(map[int]int)
	for _, c := range str {
		k := 1 << (c - 'a')
		if m[k] == 1 {
			return 0
		}
		m[k]++
		res += uint32(k)
	}
	return res
}

func max(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

package p2452

import (
	"math/bits"
)

const maskBitCompare = uint64(0b0000000010000100001000010000100001000010000100001000010000100001)
const bitOfSet = 5 // 16<26<32 =>5
const setCount = 64 / bitOfSet
const maxDiffCharCount = 2

func stringToUint64(s string, l int) []uint64 {
	if l == 0 {
		return nil
	}
	size := (l-1)/setCount + 1
	res := make([]uint64, size)
	for i := 0; i < l; i++ {
		res[i/setCount] |= uint64(s[i]-'a') << (i % setCount * bitOfSet)
	}
	return res
}

func twoEditWords(queries []string, dictionary []string) []string {
	nq, nd := len(queries), len(dictionary)
	ql, dl := len(queries[0]), len(dictionary[0])

	uQueries := make([][]uint64, nq)
	for i := 0; i < nq; i++ {
		uQueries[i] = stringToUint64(queries[i], ql)
	}

	uDictionary := make([][]uint64, nd)
	for i := 0; i < nd; i++ {
		uDictionary[i] = stringToUint64(dictionary[i], dl)
	}

	res := make([]string, 0, nq)
	for i := 0; i < nq; i++ {
		for j := 0; j < nd; j++ {
			if compare(uQueries[i], uDictionary[j]) {
				res = append(res, queries[i])
				break
			}
		}
	}

	return res
}

func compare(a, b []uint64) bool {
	diff := [maxDiffCharCount]int{}
	diffIndex := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			if diffIndex == maxDiffCharCount {
				return false
			}
			diff[diffIndex] = i
			diffIndex++
		}
	}

	count := 0
	for j := 0; j < diffIndex; j++ {
		x := a[diff[j]] ^ b[diff[j]]
		x |= (x >> 1)
		x |= (x >> 1)
		x |= (x >> 1)
		x |= (x >> 1)
		x &= maskBitCompare
		count += bits.OnesCount64(x)
	}

	return count < 3
}

//func getUint64(p *string) uint64 {
//	return binary.LittleEndian.Uint64((*[8]byte)(unsafe.Pointer(p))[:])
//}

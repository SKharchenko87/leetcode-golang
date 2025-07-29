package p2411

import (
	"math/bits"
	"slices"
)

func smallestSubarrays(nums []int) []int {
	//ToDo
}

func smallestSubarrays1(nums []int) []int {

	l := len(nums)
	maxOR := make([]int, l)
	maxOR[l-1] = nums[l-1]
	for i := l - 2; i >= 0; i-- {
		maxOR[i] = maxOR[i+1] | nums[i]
	}
	lenBitMaxOR := bits.Len(uint(maxOR[0]))
	arrIndexBits := make([]int, lenBitMaxOR)
	r := 0
	result := make([]int, l)
	for i := 0; i < l; i++ {
		for {
			if checkBit(arrIndexBits, maxOR[i]) || r == l {
				break
			}
			for j := 0; j < bits.Len(uint(nums[r])); j++ {
				arrIndexBits[j] += (nums[r] >> j) & 1
			}
			r++
		}

		result[i] = max(r-i, 1)
		for j := 0; j < bits.Len(uint(nums[i])); j++ {
			arrIndexBits[j] -= (nums[i] >> j) & 1
		}

	}
	return result
}

func checkBit(arr []int, target int) bool {
	for i := 0; i < len(arr); i++ {
		if (target>>i)&1 == 1 {
			if arr[i] == 0 {
				return false
			}
		}
	}
	return true
}

func smallestSubarrays0(nums []int) []int {

	arrIndexBits := make([][]int, 32)
	for i := 0; i < 32; i++ {
		arrIndexBits[i] = make([]int, 0)
	}

	l := len(nums)
	maxOR := 0
	for i := 0; i < l; i++ {
		maxOR |= nums[i]
		lenBit := bits.Len(uint(nums[i]))
		for j := 0; j < lenBit; j++ {
			if (nums[i]>>j)&1 == 1 {
				arrIndexBits[j] = append(arrIndexBits[j], i)
			}
		}
	}
	lenBitMaxOR := bits.Len(uint(maxOR))
	arrIndexBits = arrIndexBits[:lenBitMaxOR]
	arrBitMaxOR := make([]int, 0, lenBitMaxOR)
	for i := 0; i < lenBitMaxOR; i++ {
		if (maxOR>>i)&1 == 1 {
			arrBitMaxOR = append(arrBitMaxOR, i)
		}
	}

	result := make([]int, l)
	for i := 0; i < l; i++ {
		maxIndex := i
		for _, bitNum := range arrBitMaxOR {
			index, exists := slices.BinarySearch(arrIndexBits[bitNum], i)
			if exists {
				maxIndex = max(maxIndex, i)
			} else if index == len(arrIndexBits[bitNum]) {
				maxIndex = max(maxIndex, i)
				//break
			} else {
				maxIndex = max(maxIndex, arrIndexBits[bitNum][index])
			}
		}
		result[i] = maxIndex - i + 1
	}

	return result

}

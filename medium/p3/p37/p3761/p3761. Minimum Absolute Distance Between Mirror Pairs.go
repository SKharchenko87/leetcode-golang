package p3761

import (
	"math"
)

func minMirrorPairDistance(nums []int) int {
	lastRevert := make(map[int]int, min(1e6, len(nums)))
	res := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		if j, ok := lastRevert[nums[i]]; ok {
			res = min(res, i-j)
		}
		rev := reverse(nums[i])
		lastRevert[rev] = i + 1
	}
	if res == math.MaxInt32 {
		return -1
	}
	return res + 1
}

func reverse(num int) int {
	res := 0
	for num != 0 {
		res = (res * 10) + num%10
		num /= 10
	}
	return res
}

/*
var lastRevert map[int]int

func init() {
	lastRevert = make(map[int]int, 1e6)
}

func minMirrorPairDistance(nums []int) int {
	clear(lastRevert)
	res := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		if j, ok := lastRevert[nums[i]]; ok {
			res = min(res, i-j)
		}
		rev := reverse(nums[i])
		lastRevert[rev] = i + 1
	}
	if res == math.MaxInt32 {
		return -1
	}
	return res + 1
}

func reverse(num int) int {
	res := 0
	for num != 0 {
		res = (res * 10) + num%10
		num /= 10
	}
	return res
}
*/

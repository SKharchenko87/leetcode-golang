package p0462

import "sort"

func minMoves2(nums []int) int {
	l := len(nums)
	if l < 2 {
		return 0
	}
	sumLeft, sumRight := 0, 0
	sort.Ints(nums)
	i := 0
	for ; i < l/2; i++ {
		sumLeft += nums[i]
	}
	for ; i < l; i++ {
		sumRight += nums[i]
	}
	//equal0 := nums[l/2-1] * l / 2
	lenLeft, lenRight := l/2, l-l/2
	left, right := nums[l/2-1], nums[l/2]
	diffLeft0 := abs(sumLeft - left*lenLeft)
	diffRight0 := abs(sumRight - left*lenRight)

	//equal1 := nums[l/2] * (l - l/2)
	diffLeft1 := abs(sumLeft - right*lenLeft)
	diffRight1 := abs(sumRight - right*lenRight)

	return min(diffLeft0+diffRight0, diffLeft1+diffRight1)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

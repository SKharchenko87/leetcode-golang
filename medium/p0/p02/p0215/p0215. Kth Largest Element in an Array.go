package p0215

import (
	"math"
)

//	func findKthLargest(nums []int, k int) int {
//		sort.Ints(nums)
//		return nums[len(nums)-k]
//	}

func findKthLargest(nums []int, k int) int {
	l := len(nums)

	// Находими минимум и максимум
	minNums, maxNums := math.MaxInt, math.MinInt
	for i := 0; i < l; i++ {
		if maxNums < nums[i] {
			maxNums = nums[i]
		}
		if minNums > nums[i] {
			minNums = nums[i]
		}
	}

	// Создаем массив в котором будем хранить
	//количество чисел от  minNums-minNums до maxNums-minNums
	arr := make([]int, maxNums-minNums+1)
	for i := 0; i < l; i++ {
		arr[nums[i]-minNums]++
	}

	//Перебираем значения от maxNums до minNums
	for i := len(arr) - 1; i >= 0; i-- {
		k -= arr[i]
		if k <= 0 {
			return i + minNums
		}
	}
	return -1
}

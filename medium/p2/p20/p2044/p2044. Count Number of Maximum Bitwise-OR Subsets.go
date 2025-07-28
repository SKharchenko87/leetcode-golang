package p2044

import (
	"sort"
)

// ToDo еще есть варианты через Memorization и Bit Manipulation + Dynamic Programming

func countMaxOrSubsets(nums []int) int {
	maxOR := 0
	l := len(nums)
	for i := 0; i < l; i++ {
		maxOR |= nums[i]
	}

	var backTrack func(index int, current int) int
	backTrack = func(index int, current int) int {
		// Если достигли максимального значения, то все остальные варианты будут равны maxOR
		if current == maxOR {
			return 1 << (l - 1 - index)
		}
		index++
		if index == l {
			return 0
		}
		// С учетом значения из nums[nextIndex] + без учета nums[nextIndex]
		return backTrack(index, current|nums[index]) + backTrack(index, current)
	}

	return backTrack(0, nums[0]) + backTrack(0, 0)
}

func countMaxOrSubsets3(nums []int) int {
	maxOR := 0
	l := len(nums)
	for i := 0; i < l; i++ {
		maxOR |= nums[i]
	}

	sort.Ints(nums)

	var backTrack func(index int, isIncluded bool, current int) int
	backTrack = func(index int, isIncluded bool, current int) int {

		if isIncluded {
			current |= nums[index]
		}
		if index == 0 {
			if current == maxOR {
				return 1
			}
			return 0
		}
		return backTrack(index-1, true, current) + backTrack(index-1, false, current)
	}

	return backTrack(l-1, true, 0) + backTrack(l-1, false, 0)
}

func countMaxOrSubsets2(nums []int) int {
	maxOR := 0
	l := len(nums)
	for i := 0; i < l; i++ {
		maxOR |= nums[i]
	}
	result := 1
	// Полный перебор возможных вариантов (позиция бита числа указывает используется ли элемент из nums)
	for i := 1; i < 1<<l-1; i++ {
		currenOR := 0
		for indexBit := 0; indexBit < l; indexBit++ {
			// Если в варианте используется nums[indexBit], то учитываем его значение currentOR
			if (i>>indexBit)&1 == 1 {
				currenOR |= nums[indexBit]
			}
			// Как только текущее значение достигнет максимально добавляем этот вариант в результат
			if currenOR == maxOR {
				result++
				break
			}
		}
	}
	return result
}

func countMaxOrSubsets1(nums []int) int {
	l := len(nums)
	sum := 0
	for i := 0; i < l; i++ {
		sum |= nums[i]
	}
	res := 1
	for i := 1; i < 1<<l-1; i++ {
		x := 0
		for j := 0; j < l; j++ {
			if (i>>j)&1 == 1 {
				x |= nums[l-1-j]
			}
			if x == sum {
				res++
				break
			}
		}
	}
	return res
}

func countMaxOrSubsets0(nums []int) int {
	l := len(nums)
	sum := 0
	for i := 0; i < l; i++ {
		sum |= nums[i]
	}
	res := 0
	for i := 1; i < 1<<l; i++ {
		n := i
		x := 0
		j := l - 1
		for n > 0 {
			f := n % 2
			n /= 2
			if f == 1 {
				x |= nums[j]
			}
			if x == sum {
				res++
				break
			}
			j--
		}

	}
	return res
}

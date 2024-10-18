package p0670

import "slices"

var pow10 = []int{1, 10, 100, 1_000, 10_000, 100_000, 1_000_000, 10_000_000, 100_000_000, 100_000_000}

func getNumber(num, i int) int {
	return num / pow10[i] % 10
}

func setNumber(num, i, x int) int {
	return (num/pow10[i+1]*10+x)*pow10[i] + num%pow10[i]
}

func maximumSwap(num int) int {
	l, _ := slices.BinarySearch(pow10, num)
	maxIndex := 0
	maxValue := -1
	minIndex := -1
	for i := 0; i < l; i++ {
		if getNumber(num, i) > maxValue {
			tmp := -1
			for j := i + 1; j < l; j++ {
				if getNumber(num, i) > getNumber(num, j) {
					tmp = j
				}
			}
			if tmp != -1 {
				maxIndex = i
				maxValue = getNumber(num, i)
				minIndex = tmp
			}
		}
	}
	if minIndex == -1 {
		return num
	}
	minNumber := getNumber(num, minIndex)
	maxNumber := getNumber(num, maxIndex)
	num = setNumber(num, minIndex, maxNumber)
	num = setNumber(num, maxIndex, minNumber)
	return num
}

func maximumSwap0(num int) int {
	if num == 0 {
		return 0
	}
	l, _ := slices.BinarySearch(pow10, num)
	arr := make([]int, l)
	n := num
	for i := 0; i < l; i++ {
		arr[i] = n % 10
		n /= 10
	}
	maxIndex := 0
	maxValue := -1
	minIndex := -1
	for i := 0; i < l; i++ {
		if arr[i] > maxValue {
			tmp := -1
			for j := i + 1; j < l; j++ {
				if arr[i] > arr[j] {
					tmp = j
				}
			}
			if tmp != -1 {
				maxIndex = i
				maxValue = arr[i]
				minIndex = tmp
			}
		}
	}
	if minIndex == -1 {
		return num
	}
	arr[minIndex], arr[maxIndex] = arr[maxIndex], arr[minIndex]
	res := 0
	for i := l - 1; i >= 0; i-- {
		res = res*10 + arr[i]
	}
	return res
}

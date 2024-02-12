package p0169

import "sort"

/*
На основе того что в массиве есть всегда элемент,
который встречается более чем n/2 раз.
Например,
2 2 2 3 3 3 3
count от 2 будет расти пока не встретиться 3,
дальше он будет уменьшаться пока не станет нулевым,
а вслед мы поменяем major на 3

2 3 2 3 2 3 3
count от 2 будет колебаться от 0 до 1 и в конце будет count(3)=1
*/
func majorityElement(nums []int) (major int) {
	count := 0
	for _, v := range nums {
		if count == 0 {
			major = v
			count++
		} else if major == v {
			count++
		} else {
			count--
		}
	}
	return major
}

func majorityElement2(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	prevIndex := 0
	for i, v := range nums {
		if i > 0 && nums[i-1] != nums[i] {
			prevIndex = i
		}
		if i-prevIndex+1 > n/2 {
			return v
		}
	}
	return 0
}

func majorityElement1(nums []int) int {
	n := len(nums)
	m := make(map[int]int, 1)
	for _, v := range nums {
		m[v]++
		if m[v] > n/2 {
			return v
		}
	}
	return 0
}

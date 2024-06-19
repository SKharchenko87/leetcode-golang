package p1482

import (
	"math"
	"sort"
)

func minDays(bloomDay []int, m int, k int) int {
	n := len(bloomDay)
	if m*k > n {
		return -1
	}

	leftIndex := 0
	rightIndex := 10_000_000_000
	res := 0
loop:
	for leftIndex <= rightIndex {
		mid := (rightIndex + leftIndex) / 2
		cntBouquet := 0
		cntFlower := 0
		for i := 0; i < n; i++ {
			if bloomDay[i] < mid {
				cntFlower++
			} else {
				cntFlower = 0
			}
			if cntFlower == k {
				cntBouquet++
				cntFlower = 0
				if cntBouquet == m {
					res = mid
					rightIndex = mid - 1
					continue loop
				}
			}
		}
		leftIndex = mid + 1
	}
	return res - 1
}

func minDays1(bloomDay []int, m int, k int) int {
	n := len(bloomDay)
	if m*k > n {
		return -1
	}
	bouquets := make([]int, n-k+1)
	keys := make([]int, n)
	copy(keys, bloomDay)
	minDay := math.MaxInt
	maxDay := math.MinInt
	maxDayIndex := -1
	for i := 0; i < n-k+1; i++ {

		if maxDayIndex == i-1 {
			maxDay = math.MinInt
			for j := 0; j < k; j++ {
				if bloomDay[i+j] >= maxDay {
					maxDay = bloomDay[i+j]
					maxDayIndex = i + j
				}
			}
		} else {
			if bloomDay[i+k-1] >= maxDay {
				maxDay = bloomDay[i+k-1]
				maxDayIndex = i + k - 1
			}
		}

		bouquets[i] = maxDay
		if minDay > bouquets[i] {
			minDay = bouquets[i]
		}
	}
	sort.Ints(keys)

	leftIndex := 0
	rightIndex := n
	res := -1
	for leftIndex < rightIndex-1 {
		mid := (rightIndex + leftIndex) / 2
		key := keys[mid]
		flg := false
		if key >= minDay {
			cnt := 0
			for i := 0; i < len(bouquets); i++ {
				if bouquets[i] <= key {
					cnt++
					i = i + k - 1
				}
				if cnt == m {
					flg = true
					res = mid
					break
				}
			}
		}
		if flg {
			rightIndex = mid
		} else {
			leftIndex = mid
		}

	}

	return keys[res]
}

func minDays0(bloomDay []int, m int, k int) int {
	n := len(bloomDay)
	if m*k > n {
		return -1
	}
	bouquets := make([]int, n-k+1)
	md := make(map[int]int, n)
	keys := []int{}
	minDay := math.MaxInt
	for i := 0; i < n-k+1; i++ {

		maxDay := math.MinInt
		for j := 0; j < k; j++ {
			if _, exist := md[bloomDay[i+j]]; !exist {
				md[bloomDay[i+j]] = 0
				keys = append(keys, bloomDay[i+j])
			}
			if bloomDay[i+j] > maxDay {
				maxDay = bloomDay[i+j]
			}
		}
		bouquets[i] = maxDay
		if minDay > bouquets[i] {
			minDay = bouquets[i]
		}
	}
	sort.Ints(keys)

	for _, key := range keys {
		if key >= minDay {
			cnt := 0
			for i := 0; i < len(bouquets); i++ {
				if bouquets[i] <= key {
					cnt++
					i = i + k - 1
				}
				if cnt == m {
					return key
				}
			}
		}
	}

	return -1
}

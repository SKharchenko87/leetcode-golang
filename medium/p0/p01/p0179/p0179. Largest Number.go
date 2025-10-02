package p0179

import (
	"slices"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	slices.SortFunc(nums, func(a, b int) int {
		if a*b == 0 {
			return b - a
		}
		return b*(lenNum(a)-1) - a*(lenNum(b)-1)
	})
	if nums[0] == 0 {
		return "0"
	}
	b := strings.Builder{}
	for i := 0; i < len(nums); i++ {
		b.WriteString(strconv.Itoa(nums[i]))
	}
	return b.String()
}

var pow10 = []int{1, 10, 100, 1_000, 10_000, 100_000, 1_000_000, 10_000_000, 100_000_000, 1_000_000_000, 10_000_000_000}

func lenNum(num int) int {
	index, flg := slices.BinarySearch(pow10, num)
	if flg {
		return pow10[index+1]
	}
	return pow10[index]
}

func lenNum0(num int) int {
	p := 1
	for num/p > 0 {
		p *= 10
	}
	return p
}

func largestNumber1(nums []int) string {
	strs := make([]string, len(nums))
	for i, num := range nums {
		strs[i] = strconv.Itoa(num)
	}
	slices.SortFunc(strs, func(a, b string) int {
		return -strings.Compare(a+b, b+a)
	})
	if strs[0] == "0" {
		return "0"
	}
	return strings.Join(strs, "")
}

func largestNumber0(nums []int) string {
	b := strings.Builder{}
	slices.SortFunc(nums, func(a, b int) int {
		slA, slB := numToSlice(a), numToSlice(b)
		lA, lB := len(slA), len(slB)
		for i := 0; i < lA+lB; i++ {
			var ab, ba int
			if i < lA {
				ab = slA[i]
			} else {
				ab = slB[i%lA]
			}
			if i < lB {
				ba = slB[i]
			} else {
				ba = slA[i%lB]
			}
			if r := ba - ab; r != 0 {
				return r
			}
		}
		return 0
	})
	if nums[0] == 0 {
		return "0"
	}
	for i := 0; i < len(nums); i++ {
		b.WriteString(strconv.Itoa(nums[i]))
	}
	return b.String()
}

func numToSlice(num int) (res []int) {
	if num == 0 {
		return []int{0}
	}
	res = make([]int, 0, 9)
	for num > 0 {
		defer func(res *[]int, i int) {
			*res = append(*res, i)
		}(&res, num%10)
		num /= 10
	}
	return
}

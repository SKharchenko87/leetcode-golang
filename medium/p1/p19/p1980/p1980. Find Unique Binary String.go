package p1980

import (
	"fmt"
	"strconv"
)

func findDifferentBinaryString(nums []string) string {
	n := len(nums)
	res := make([]byte, n)
	for i := 0; i < n; i++ {
		res[i] = '0'
	}
	for i := 0; i < n; i++ {
		if nums[i][i] == '0' {
			res[i] = '1'
		} else {
			res[i] = '0'
		}
	}
	return string(res)
}

func findDifferentBinaryString1(nums []string) string {
	n := len(nums)
	numbers := make([]uint, n)
	for i := 0; i < n; i++ {
		x, _ := strconv.ParseUint(nums[i], 2, 16)
		numbers[i] = uint(x)
	}
	for i := uint(0); i < 1<<n; i++ {
		if !check(numbers, i) {
			return fmt.Sprintf("%016b", i)[16-n:]
		}
	}
	return ""
}

func check(numbers []uint, x uint) bool {
	for _, number := range numbers {
		if x == number {
			return true
		}
	}
	return false
}

type Tree struct {
	children [2]*Tree
}

func findDifferentBinaryString0(nums []string) string {
	n := byte(len(nums))
	tree := Tree{}

	for _, num := range nums {
		root := &tree
		for i := byte(0); i < n; i++ {
			var bit byte
			if num[i] == '0' {
				bit = 0
			} else {
				bit = 1
			}
			if root.children[bit] == nil {
				root.children[bit] = &Tree{}
			}
			root = root.children[bit]
		}
	}
	res := make([]byte, n)
	for i := byte(0); i < n; i++ {
		res[i] = '0'
	}
	var dfs func(root *Tree, level byte) bool
	dfs = func(root *Tree, level byte) bool {
		if level == n {
			return false
		}
		res[level] = '0'
		if root.children[0] == nil {
			return true
		}
		if dfs(root.children[0], level+1) {
			return true
		}

		res[level] = '1'
		if root.children[1] == nil {
			res[level] = '1'
			return true
		}
		if dfs(root.children[1], level+1) {
			return true
		}
		res[level] = '0'
		return false
	}

	dfs(&tree, 0)
	return string(res)
}

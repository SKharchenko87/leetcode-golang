package p0703

import (
	"slices"
	"sort"
)

type KthLargest struct {
	k    int
	nums []int
}

func Constructor(k int, nums []int) KthLargest {
	result := KthLargest{k: k, nums: make([]int, 0, k)}
	for _, num := range nums {
		result.Add(num)
	}
	return result
}

func (this *KthLargest) Add(val int) int {
	l := len(this.nums)
	if l == 0 {
		this.nums = append(this.nums, val)
	} else if l < this.k {
		index := len(this.nums) - 1
		this.nums = append(this.nums, this.nums[index])
		for ; index >= 0 && this.nums[index] > val; index-- {
			this.nums[index+1] = this.nums[index]
		}
		this.nums[index+1] = val

	} else if this.nums[0] < val {
		index := 0
		for ; this.nums[index] < val; index++ {
			if index < l-1 {
				this.nums[index] = this.nums[index+1]
			} else {
				index++
				break
			}
		}
		this.nums[index-1] = val
	}
	return this.nums[0]
}

func Constructor1(k int, nums []int) KthLargest {
	l := len(nums)
	if l < k {
		nums = append(nums, -10001)
		l++
	}
	sort.Ints(nums)
	nums = nums[l-k:]
	return KthLargest{k: k, nums: nums}
}

func (this *KthLargest) Add1(val int) int {
	if this.nums[0] < val {
		l := len(this.nums)
		index := 0
		for ; this.nums[index] < val; index++ {
			if index < l-1 {
				this.nums[index] = this.nums[index+1]
			} else {
				index++
				break
			}
		}
		this.nums[index-1] = val
	}
	return this.nums[0]
}

func run(commands []string, k int, initParams []int, params []int) []int {
	var obj KthLargest
	obj = Constructor(k, initParams)
	var res []int
	for i := 0; i < len(commands); i++ {
		res = append(res, obj.Add(params[i]))
	}
	return res
}

func Constructor0(k int, nums []int) KthLargest {
	sort.Ints(nums)
	return KthLargest{k: k, nums: nums}
}

func (this *KthLargest) Add0(val int) int {
	index, _ := slices.BinarySearch(this.nums, val)
	this.nums = append(this.nums[:index], append([]int{val}, this.nums[index:]...)...)
	return this.nums[len(this.nums)-this.k]
}

package p0962

func maxWidthRamp(nums []int) int {
	l := len(nums)
	index := 0
	stack := make([]int, l)
	stack[0] = 0
	for i := 1; i < l; i++ {
		if nums[stack[index]] > nums[i] {
			index++
			stack[index] = i
		}
	}
	res := 0
	for j := l - 1; j >= 0; j-- {
		for index >= 0 && nums[stack[index]] <= nums[j] {
			res = max(res, j-stack[index])
			index--
		}
	}
	return res
}

func maxWidthRamp2(nums []int) int {
	l := len(nums)
	index := l - 1
	stack := make([]int, l)
	stack[index] = index
	for i := l - 2; i >= 0; i-- {
		if nums[i] > nums[stack[index]] {
			index--
			stack[index] = i
		}
	}
	res := 0
	for i := 0; i < l-res; i++ {
		for ; index < l && nums[i] <= nums[stack[index]]; index++ {
			res = max(res, stack[index]-i)
		}
	}
	return res
}

func maxWidthRamp1(nums []int) int {
	l := len(nums)
	index := l - 1
	stack := make([]int, l)
	stack[index] = index
	for i := l - 2; i >= 0; i-- {
		if nums[i] > nums[stack[index]] {
			index--
			stack[index] = i
		}
	}
	res := 0
	for i := 0; i < l-res; i++ {
		for j := l - 1; j >= index; j-- {
			if nums[i] <= nums[stack[j]] {
				if res < stack[j]-i {
					res = stack[j] - i
				}
				break
			}
		}
	}
	return res
}

// TLE
func maxWidthRamp0(nums []int) int {
	res := 0
	for i := 0; i < len(nums)-res; i++ {
		for j := len(nums) - 1; j-i >= res; j-- {
			if nums[i] <= nums[j] {
				if res < j-i {
					res = j - i
				}
				break
			}
		}
	}
	return res
}

// 9,8,1,0,1,9,4,0,4,1
// _,_,_,_,_,9,_,_,4,1

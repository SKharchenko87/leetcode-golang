package p2044

func countMaxOrSubsets(nums []int) int {
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

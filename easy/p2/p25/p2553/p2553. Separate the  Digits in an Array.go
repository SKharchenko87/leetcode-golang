package p2553

func rec(res *[]int, n int) {
	if n == 0 {
		return
	}
	d := n % 10
	rec(res, n/10)
	*res = append(*res, d)
}

func separateDigits(nums []int) []int {
	l := len(nums)
	res := make([]int, 0, 10*l)
	for i := 0; i < l; i++ {
		rec(&res, nums[i])
	}
	return res
}

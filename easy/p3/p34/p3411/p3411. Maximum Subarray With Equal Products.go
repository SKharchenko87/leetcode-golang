package p3411

func maxLength(nums []int) int {
	res := 0
	for i := 0; i < len(nums)-1; i++ {

		_gcd := nums[i]
		lcm := nums[i]
		p := nums[i]
		j := i + 1
		for ; j < len(nums); j++ {
			_gcd = gcd(nums[j], _gcd)
			lcm = lcm * nums[j] / gcd(lcm, nums[j])
			p *= nums[j]
			if p != _gcd*lcm {
				break
			}
		}
		res = max(res, j-i)
	}
	return res
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

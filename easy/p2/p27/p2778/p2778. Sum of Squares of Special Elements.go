package p2778

var m = [50][]int{}

func init() {
	for i := 0; i < 50; i++ {
		m[i] = make([]int, 0, 6)
	}
	for i := 0; i < 50; i++ {
		for j := 0; j <= i; j++ {
			if (i+1)%(j+1) == 0 {
				m[i] = append(m[i], j)
			}
		}
	}
}

func sumOfSquares(nums []int) int {
	n := len(nums)
	res := 0
	for _, index := range m[n-1] {
		res += nums[index] * nums[index]
	}
	return res
}

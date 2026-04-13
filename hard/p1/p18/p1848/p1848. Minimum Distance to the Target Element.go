package p1848

func getMinDistance(nums []int, target int, start int) int {
	n := len(nums)
	res := n + 1
	for i := 0; i < n; i++ {
		if nums[i] == target {
			if x := abs(i - start); x < res {
				res = x
				n = min(n, x+start)
			}
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

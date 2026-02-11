package p3721

func longestBalanced(nums []int) int {
	n := len(nums)
	lastIndex := make(map[int]int, n)
	prevIndex := make([]int, n)
	totalOdd, totalEven := 0, 0
	for i := 0; i < n; i++ {
		if v, ok := lastIndex[nums[i]]; !ok {
			prevIndex[i] = -1
			if nums[i]%2 == 0 {
				totalEven++
			} else {
				totalOdd++
			}
		} else {
			prevIndex[i] = v
		}
		lastIndex[nums[i]] = i
	}

	res := 0
	for i := 0; i < n && n-i > res; i++ {

		curTotalEven := totalEven
		curTotalOdd := totalOdd
		for j := n - 1; j > i && res < j-i+1 && curTotalOdd > 0 && curTotalEven > 0; j-- {
			if curTotalOdd == curTotalEven {
				res = j - i + 1
				break
			}
			if prevIndex[j] < i {
				if nums[j]%2 == 0 {
					curTotalEven--
				} else {
					curTotalOdd--
				}
			}
		}

		if lastIndex[nums[i]] == i {
			if nums[i]%2 == 0 {
				totalEven--
			} else {
				totalOdd--
			}
		}
	}
	return res
}

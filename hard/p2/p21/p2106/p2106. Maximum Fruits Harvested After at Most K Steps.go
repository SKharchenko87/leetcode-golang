package p2106

func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	length := max(fruits[len(fruits)-1][0], startPos+k) + 1
	preSum := make([]int, length)
	prev := 0
	index := 0
	for i := 0; i < length; i++ {
		preSum[i] += prev
		if index < len(fruits) && i == fruits[index][0] {
			preSum[i] += fruits[index][1]
			index++
		}
		prev = preSum[i]
	}

	getRangeSum := func(l, r int) int {
		if l > r {
			return 0
		}
		if l == 0 {
			return preSum[r]
		}
		return preSum[r] - preSum[l-1]
	}

	res := 0

	// Сначала направо, потом налево
	for x := 0; x <= k; x++ {
		rightEnd := startPos + x
		leftSteps := k - 2*x
		if leftSteps < 0 {
			continue
		}
		leftStart := max(0, startPos-leftSteps)
		if rightEnd < length {
			res = max(res, getRangeSum(leftStart, rightEnd))
		}
	}

	// Сначала налево, потом направо
	for x := 0; x <= k; x++ {
		leftStart := startPos - x
		rightSteps := k - 2*x
		if rightSteps < 0 {
			continue
		}
		rightEnd := startPos + rightSteps
		if leftStart >= 0 {
			res = max(res, getRangeSum(leftStart, rightEnd))
		}
	}

	return res
}

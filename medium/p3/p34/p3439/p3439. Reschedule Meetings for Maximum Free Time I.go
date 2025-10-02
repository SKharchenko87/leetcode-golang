package p3439

func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int {
	prevStartTime := 0
	currentStartTime := 0
	i := 0
	sumWindow := 0
	for ; i < k; i++ {
		currentStartTime = startTime[i]
		sumWindow += currentStartTime - prevStartTime
		prevStartTime = endTime[i]
	}
	res := sumWindow

	if i != len(startTime) {
		currentStartTime = startTime[i]
		sumWindow += currentStartTime - prevStartTime
		prevStartTime = endTime[i]
		res = max(res, sumWindow)
		sumWindow -= startTime[i-k]
		i++

		for ; i < len(startTime); i++ {
			currentStartTime = startTime[i]
			sumWindow += currentStartTime - prevStartTime
			prevStartTime = endTime[i]

			res = max(res, sumWindow)
			sumWindow -= startTime[i-k] - endTime[i-k-1]
		}
	}

	sumWindow += eventTime - prevStartTime
	res = max(res, sumWindow)
	return res
}

func maxFreeTime3(eventTime int, k int, startTime []int, endTime []int) int {
	prevStartTime := 0
	currentStartTime := 0
	i := 0
	sumWindow := 0
	var f = func() {
		currentStartTime = startTime[i]
		sumWindow += currentStartTime - prevStartTime
		prevStartTime = endTime[i]
	}
	for ; i < k; i++ {
		f()
	}
	res := sumWindow

	if i != len(startTime) {
		f()
		res = max(res, sumWindow)
		sumWindow -= startTime[i-k]
		i++

		for ; i < len(startTime); i++ {
			f()
			res = max(res, sumWindow)
			sumWindow -= startTime[i-k] - endTime[i-k-1]
		}
	}

	sumWindow += eventTime - prevStartTime
	res = max(res, sumWindow)
	return res
}

func maxFreeTime2(eventTime int, k int, startTime []int, endTime []int) int {
	prevStartTime := 0
	currentStartTime := 0
	i := 0
	sumWindow := 0
	for ; i < k; i++ {
		currentStartTime = startTime[i]
		sumWindow += currentStartTime - prevStartTime
		prevStartTime = endTime[i]
	}

	res := 0
	for ; i < len(startTime); i++ {
		currentStartTime = startTime[i]
		sumWindow += currentStartTime - prevStartTime
		prevStartTime = endTime[i]

		res = max(res, sumWindow)
		if i-k-1 >= 0 {
			sumWindow -= startTime[i-k] - endTime[i-k-1]
		} else {
			sumWindow -= startTime[i-k]
		}

	}

	sumWindow += eventTime - prevStartTime
	res = max(res, sumWindow)
	return res
}

func maxFreeTime1(eventTime int, k int, startTime []int, endTime []int) int {
	prevStartTime := 0
	currentStartTime := 0
	cnt := 0
	sumWindow := 0
	res := 0
	for i := 0; i < len(startTime); i++ {
		currentStartTime = startTime[i]
		sumWindow += currentStartTime - prevStartTime
		prevStartTime = endTime[i]
		cnt++
		if cnt > k {
			res = max(res, sumWindow)
			if i-k-1 >= 0 {
				sumWindow -= startTime[i-k] - endTime[i-k-1]
			} else {
				sumWindow -= startTime[i-k]
			}
			cnt--
		}
	}

	currentStartTime = eventTime
	sumWindow += currentStartTime - prevStartTime
	res = max(res, sumWindow)
	return res
}

func maxFreeTime0(eventTime int, k int, startTime []int, endTime []int) int {
	freeTime := make([]int, 0, len(startTime)+1)
	prevStartTime := 0
	currentStartTime := 0
	for i := 0; i < len(startTime); i++ {
		currentStartTime = startTime[i]
		freeTime = append(freeTime, currentStartTime-prevStartTime)
		prevStartTime = endTime[i]
	}
	freeTime = append(freeTime, eventTime-prevStartTime)
	sumWindow := 0
	for i := 0; i <= k; i++ {
		sumWindow = sumWindow + freeTime[i]
	}
	res := sumWindow
	for i := k + 1; i < len(freeTime); i++ {
		sumWindow -= freeTime[i-k-1]
		sumWindow += freeTime[i]
		res = max(res, sumWindow)
	}
	return res
}

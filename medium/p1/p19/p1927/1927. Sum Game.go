package p1927

func sumGame(num string) bool {
	leftCount, rightCount := 0, 0
	leftSum, rightSum := 0, 0
	i := 0
	for ; i < len(num)/2; i++ {
		if num[i] == '?' {
			leftCount++
		} else {
			leftSum += int(num[i] - '0')
		}
	}
	for ; i < len(num); i++ {
		if num[i] == '?' {
			rightCount++
		} else {
			rightSum += int(num[i] - '0')
		}
	}
	return (leftCount+rightCount)%2 == 1 || leftSum-rightSum != (rightCount-leftCount)*9/2
}

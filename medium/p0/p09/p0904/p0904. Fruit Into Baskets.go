package p0904

func totalFruit(fruits []int) int {
	ch0 := fruits[0]
	lastCh0Index := 0
	lastCh1Index := 0
	startPos := 0
	l := len(fruits)
	for ; lastCh1Index < l && ch0 == fruits[lastCh1Index]; lastCh1Index++ {
	}

	if lastCh1Index == l {
		return l
	}

	lastCh0Index = lastCh1Index - 1
	ch1 := fruits[lastCh1Index]

	res := lastCh1Index - startPos + 1

	for i := lastCh1Index + 1; i < l; i++ {
		switch fruits[i] {
		case ch0:
			lastCh0Index = i
		case ch1:
			lastCh1Index = i
		default:
			res = max(res, i-startPos)
			if lastCh0Index < lastCh1Index {
				startPos = lastCh0Index + 1
				ch0 = ch1
				lastCh0Index = lastCh1Index
			} else {
				startPos = lastCh1Index + 1
			}
			ch1 = fruits[i]
			lastCh1Index = i
		}
	}
	res = max(res, l-startPos)
	return res
}

func totalFruit0(fruits []int) int {
	basket := make(map[int]int)
	left := 0
	maxFruits := 0

	for right, fruit := range fruits {
		basket[fruit]++

		for len(basket) > 2 {
			leftFruit := fruits[left]
			basket[leftFruit]--
			if basket[leftFruit] == 0 {
				delete(basket, leftFruit)
			}
			left++
		}

		if current := right - left + 1; current > maxFruits {
			maxFruits = current
		}
	}
	return maxFruits
}

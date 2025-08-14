package p0624

func maxDistance(arrays [][]int) int {
	min0, min1, max0, max1 := arrays[0][0], arrays[1][0], arrays[0][len(arrays[0])-1], arrays[1][len(arrays[1])-1]
	min0Index, min1Index, max0Index, max1Index := 0, 1, 0, 1
	if min0 > min1 {
		min0, min1 = min1, min0
		min0Index, min1Index = min1Index, min0Index
	}
	if max0 < max1 {
		max0, max1 = max1, max0
		max0Index, max1Index = max1Index, max0Index
	}
	for i := 2; i < len(arrays); i++ {
		if arrays[i][0] < min1 {
			if arrays[i][0] <= min0 {
				min1 = min0
				min1Index = min0Index
				min0 = arrays[i][0]
				min0Index = i
			} else {
				min1 = arrays[i][0]
				min1Index = i
			}
		}
		maxIndex := len(arrays[i]) - 1
		if arrays[i][maxIndex] > max1 {
			if arrays[i][maxIndex] >= max0 {
				max1 = max0
				max1Index = max0Index
				max0 = arrays[i][maxIndex]
				max0Index = i
			} else {
				max1 = arrays[i][maxIndex]
				max1Index = i
			}
		}
	}
	if max0Index == min0Index {
		return max(max0-min1, max1-min0)
	} else {
		return max0 - min0
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

package p0042

import "sort"

func trap(height []int) int {
	l := len(height)
	sum := 0
	arr := make([][2]int, l)
	for i := 0; i < l; i++ {
		sum += height[i]
		arr[i] = [2]int{height[i], i}
	}

	sort.Slice(arr, func(i, j int) bool {
		if arr[i][0] == arr[j][0] {
			return arr[i][1] < arr[j][1]
		} else {
			return arr[i][0] > arr[j][0]
		}
	})

	maxHeight := arr[0][0]
	maxLeftIndex := arr[0][1]
	maxRightIndex := arr[0][1]

	for ; maxRightIndex < l && height[maxRightIndex] == maxHeight; maxRightIndex++ {
	}
	maxRightIndex--

	res := 0

	for i := 1; i < maxLeftIndex; i++ {
		if height[i-1] > height[i] {
			res += height[i-1] - height[i]
			height[i] = height[i-1]
		}
	}

	for i := maxLeftIndex + 1; i < maxRightIndex; i++ {
		res += maxHeight - height[i]
	}

	for i := l - 1; i > maxRightIndex; i-- {
		if height[i] > height[i-1] {
			res += height[i] - height[i-1]
			height[i-1] = height[i]
		}
	}

	return res
}

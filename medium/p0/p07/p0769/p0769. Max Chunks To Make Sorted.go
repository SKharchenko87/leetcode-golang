package p0769

func maxChunksToSorted(arr []int) int {
	cnt := 0
	maxVal := 0
	for i := 0; i < len(arr); i++ {
		maxVal = max(maxVal, arr[i])
		if maxVal == i {
			cnt++
		}
	}
	return cnt
}

func maxChunksToSorted0(arr []int) int {
	cnt := 0
	for i := 0; i < len(arr); i++ {
		maxVal := arr[i]
		j := i
		for ; j < len(arr) && arr[j] != i; j++ {
			if arr[j] > maxVal {
				maxVal = arr[j]
			}
		}
		for ; j < len(arr) && maxVal >= j; j++ {
			if arr[j] > maxVal {
				maxVal = arr[j]
			}
			if maxVal == j {
				break
			}
		}

		i = j

		cnt++
	}
	return cnt
}

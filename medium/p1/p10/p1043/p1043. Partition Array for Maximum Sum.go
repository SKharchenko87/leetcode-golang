package p1043

import "fmt"

func maxSumAfterPartitioning(arr []int, k int) int {
	l := len(arr)
	res := make([]int, l)
	var findMax = func(j int) (int, int) {
		indexMax := j
		valueMax := arr[indexMax]
		for i := 0; i < k-1; i++ {
			curIndex := (l + j - k + i + 1) % l
			curValue := arr[curIndex]
			if valueMax < curValue {
				indexMax = curIndex
				valueMax = curValue
			}
		}
		return indexMax, valueMax
	}
	indexMax, valueMax := findMax(0)
	res[0] = valueMax
	for i := 1; i < l; i++ {
		if valueMax <= arr[i] {
			valueMax = arr[i]
			indexMax = i
		}
		if indexMax == (l+i-k)%l {
			indexMax, valueMax = findMax(i)
		}
		res[i] = valueMax
	}
	fmt.Println(res)
	return 0
}

package p0997

func findJudge(n int, trust [][]int) int {
	if n == 1 {
		return 1
	}
	i, arr, judgeIndex := 0, make([]int, n), -1
	for ; i < len(trust); i++ {
		if judgeIndex == trust[i][0] {
			return -1
		}
		arr[trust[i][0]-1] = -1
		index := trust[i][1] - 1
		if arr[index] != -1 {
			arr[index]++
			if arr[index] == n-1 {
				judgeIndex = index + 1
			}
		}
	}
	return judgeIndex
}

func findJudge1(n int, trust [][]int) int {
	i := 0
	l := len(trust)
	if l == 0 && n == 1 {
		return 1
	}
	arr := make([]int, n)
	judgeIndex := -1
	for ; i < l; i++ {
		arr[trust[i][0]-1] = -1
		index := trust[i][1] - 1
		if arr[index] != -1 {
			arr[index]++
			if arr[index] == n-1 {
				judgeIndex = index + 1
				break
			}
		}
	}
	for ; i < l; i++ {
		if judgeIndex == trust[i][0] {
			return -1
		}
	}
	return judgeIndex
}

package p1526

func minNumberOperations(target []int) int {
	res := target[0]
	for i := 1; i < len(target); i++ {
		res += max(0, target[i]-target[i-1])
	}
	return res
}

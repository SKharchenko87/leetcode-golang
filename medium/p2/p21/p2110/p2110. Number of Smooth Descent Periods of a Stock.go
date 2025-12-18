package p2110

func getDescentPeriods(prices []int) int64 {
	var result, seq int64
	seq = 1
	result = 1
	for i := 1; i < len(prices); i++ {
		if prices[i]+1 == prices[i-1] {
			seq++
		} else {
			seq = 1
		}
		result += seq
	}
	return result
}

package p3494

func minTime(skill []int, mana []int) int64 {
	n := len(skill)
	m := len(mana)
	prefSum := make([]int, n)
	prefSum[0] = skill[0]
	for i := 1; i < n; i++ {
		prefSum[i] = prefSum[i-1] + skill[i]
	}
	var res int64
	var start int64
	i := 0
	for ; i < m-1; i++ {
		newStart := start + int64(prefSum[0]*mana[i])
		res = start + int64(prefSum[0]*mana[i])
		for j := 1; j < n; j++ {
			res = start + int64(prefSum[j]*mana[i])
			newStart = max(newStart, res-int64(prefSum[j-1]*mana[i+1]))
		}

		start = newStart
	}
	res = start + int64(prefSum[0]*mana[i])
	for j := 1; j < n; j++ {
		res = start + int64(prefSum[j]*mana[i])
	}
	return res
}

/*      1    5       2     4
0	0	5	30	     40	  60
    52 = 60 - 1 - 5 -2 				 - max
    34 = 40 - 1 - 5
    29 = 30 - 1
	5 = 5
1	52	53	58	60	64
    32 = 64 - 4 - 20 - 8
    36 = 60 - 4 - 20
    54 = 58 - 4     				- max
	53
2	54	58	78	86	102
	86 = 102 -2 - 10 - 4			- max
	74 = 86 -2 - 10
	76 = 78 -2
	58
3	86	88	98	102	110
*/

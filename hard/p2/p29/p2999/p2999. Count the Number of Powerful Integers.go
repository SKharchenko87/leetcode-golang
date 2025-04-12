package p2999

import "slices"

func numberOfPowerfulInt(start int64, finish int64, limit int, s string) int64 {
	before := getNumber(start-1, limit, s)
	all := getNumber(finish, limit, s)
	return all - before
}

func getNumber(num int64, limit int, s string) int64 {
	count, _ := slices.BinarySearch(powX[10], num-1)
	count++
	l := len(s)
	var x int64
	var b bool
	for i := count - 1; i >= l; i-- {
		d := getDigit(num, i)
		b = b || (d > limit)
		if b {
			x += int64(limit) * powX[10][i]
		} else {
			x += int64(d) * powX[10][i]
		}
	}
	if b {
		for i := 0; i < l; i++ {
			x += int64(s[l-1-i]-'0') * powX[10][i]
		}
	} else {
		x += num % powX[10][l]
	}

	var res int64 = 0
	for i := l; i < count; i++ {
		d := getDigit(x, i)
		dm := min(d, limit)
		res += int64(dm) * powX[limit+1][i-l]
	}
	if comparePostfix(x, s) >= 0 {
		res++
	}
	return res
}

func comparePostfix(num int64, s string) int {
	l := len(s)
	for i := l - 1; i >= 0; i-- {
		v := getDigit(num, i) - int(s[l-1-i]-'0')
		if v == 0 {
			continue
		} else if v < 0 {
			return -1
		} else {
			return 1
		}
	}
	return 0
}

const MaxRankNum = 16
const MaxDigital = 10

var powX = make([][]int64, MaxDigital+1)

func init() {
	for i := 1; i < len(powX); i++ {
		powX[i] = make([]int64, MaxRankNum+1)
		powX[i][0] = 1
		for j := 1; j < len(powX[i]); j++ {
			powX[i][j] = powX[i][j-1] * int64(i)
		}
	}
}

func getDigit(num int64, index int) int {
	return int(num / powX[10][index] % 10)
}

package p2264

var res = []string{"", "000", "111", "222", "333", "444", "555", "666", "777", "888", "999"}

func largestGoodInteger(num string) string {
	var resIndex uint8 = 0
	i := 2
	for ; i < len(num); i++ {
		if num[i] == num[i-1] && num[i] == num[i-2] && resIndex < num[i]-'0'+1 {
			resIndex = num[i] - '0' + 1
		}
	}
	return res[resIndex]
}

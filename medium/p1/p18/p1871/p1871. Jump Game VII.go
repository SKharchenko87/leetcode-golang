package p1871

//ToDo
/*
func canReach(s string, minJump, maxJump int) bool {
	l := len(s)
	if s[0] != '0' && s[l-1] != '0' {
		return false
	}
	cnts := make([]int, l)
	prevZero := make([]int, l)
	for i := 1; i < l; i++ {
		cnts[i] = cnts[i-1]
		prevZero[i] = prevZero[i-1]
		if s[i] == '0' {
			cnts[i]++
			prevZero[i] = i
		}
	}

	first, last := 0, 0

	for i, index := range indexes {
		firstTmp, _ := slices.BinarySearch(indexes, first+minJump)
		lastTmp, _ := slices.BinarySearch(indexes, last+maxJump)
		if firstTmp > index+maxJump {

		}

	}
}
*/

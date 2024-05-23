package p0131

func partition(s string) (res [][]string) {
	l := len(s)
	fragments := map[[2]int]bool{}
	var isPalindrome = func(startIndex, endIndex int) bool {
		index := [2]int{startIndex, endIndex}
		if v, ok := fragments[index]; ok {
			return v
		}
		for i := 0; i < (endIndex-startIndex)/2+1; i++ {
			if s[startIndex+i] != s[endIndex-i] {
				fragments[index] = false
				return false
			}
		}
		fragments[index] = true
		return true
	}
	var backTracking func(candidate []string, prevIndex, index int)
	backTracking = func(candidate []string, prevIndex, index int) {
		if index == l {
			if prevIndex == index {
				res = append(res, candidate)
			}
			return
		}
		newCandidate := make([]string, len(candidate))
		copy(newCandidate, candidate)
		if isPalindrome(prevIndex, index) {
			candidate = append(candidate, s[prevIndex:index+1])
			backTracking(candidate, index+1, index+1)
		}
		backTracking(newCandidate, prevIndex, index+1)
	}
	backTracking([]string{}, 0, 0)
	return res
}

func partition0(s string) (res [][]string) {
	l := len(s)
	b := []byte(s)
	bad := map[int]int{}
	good := map[int]int{}
loopExternal:
	for i := 0; i < 1<<l; i++ {
		tmp := []string{}
		ts := 0
		cnt := 0
		for j := 0; j < l; j++ {
			if i>>j&1 == 0 {
				ch := bad[ts*100+j]
				ch2 := good[ts*100+j]
				if ch == 0 && (ch2 > 0 || checkPalindrome(s, ts, j)) {
					good[ts*100+j]++
					tmp = append(tmp, string(b[ts:j+1]))
					cnt += j + 1 - ts
				} else {
					bad[ts*100+j]++
					continue loopExternal
				}
				//println(i, ts, j)
				ts = j + 1
			}
		}
		if cnt == l {
			res = append(res, tmp)
		}
	}
	return res
}

func checkPalindrome(s string, startIndex, endIndex int) bool {
	for i := 0; i < (endIndex-startIndex)/2+1; i++ {
		if s[startIndex+i] != s[endIndex-i] {
			return false
		}
	}
	return true
}

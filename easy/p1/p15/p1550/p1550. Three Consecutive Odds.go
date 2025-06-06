package p1550

func threeConsecutiveOdds(arr []int) bool {
	for i := 0; i < len(arr)-2; i++ {
		if arr[i]&1+arr[i+1]&1+arr[i+2]&1 == 3 {
			return true
		}
	}
	return false
}

func threeConsecutiveOdds1(arr []int) bool {
	cntOdd := 0
	for _, x := range arr {
		if x%2 == 1 {
			if cntOdd == 2 {
				return true
			}
			cntOdd++
		} else {
			cntOdd = 0
		}
	}
	return false
}

func threeConsecutiveOdds0(arr []int) bool {
	cnt := 0
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			cnt = 0
		} else {
			cnt++
			if cnt >= 3 {
				return true
			}
		}
	}
	return false
}

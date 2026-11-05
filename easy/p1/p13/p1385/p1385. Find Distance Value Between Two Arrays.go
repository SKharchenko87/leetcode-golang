package p1385

func findTheDistanceValue(arr1, arr2 []int, d int) int {
	cnt := 0
LOOP:
	for _, n1 := range arr1 {
		for _, n2 := range arr2 {
			if abs(n1-n2) <= d {
				continue LOOP
			}
		}
		cnt++
	}
	return cnt
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

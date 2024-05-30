package p1442

func countTriplets(arr []int) int {
	count := 0
	l := len(arr)
	for i := 0; i < l-1; i++ {
		x := arr[i]
		for j := i + 1; j < l; j++ {
			x ^= arr[j]
			if x == 0 {
				count += j - i
			}
		}
	}
	//for i, x := range arr {
	//	for j, y := range arr[i+1:] {
	//		x ^= y
	//		if x == 0 {
	//			count += j + 1
	//		}
	//	}
	//}
	return count
}

func countTriplets0(arr []int) int {
	l := len(arr)
	var a, b, res int
	for i := 0; i < l-1; i++ {
		a = arr[i]
		for j := i + 1; j < l; j++ {
			a ^= arr[j]
			b = arr[j]
			for k := j; k < l; k++ {
				b ^= arr[k]
				//println(i, j, k, a, k)
				if a == b {
					res++
				}
			}

		}
	}
	return res
}

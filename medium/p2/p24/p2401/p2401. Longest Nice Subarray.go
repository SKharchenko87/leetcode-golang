package p2401

func longestNiceSubarray(nums []int) int {
	//for _, num := range nums {
	//	fmt.Printf("%32b %d\n", num, num)
	//}
	bits := [32]int{}
	lastIndex := [32]int{}
	for i := 0; i < 32; i++ {
		bits[i] = 1 << i
		lastIndex[i] = -1
	}
	res := 0
	prevDistance := 1

	for i, num := range nums {
		minDistance := 1_000_000_000
		for indexBit, bit := range bits {
			if num&bit == bit {
				if minDistance > i-lastIndex[indexBit] {
					minDistance = i - lastIndex[indexBit]
				}
				lastIndex[indexBit] = i
			}
		}
		if i == 0 {
			minDistance = 1
		}
		if minDistance > prevDistance {
			minDistance = prevDistance + 1
		}
		if minDistance > res {
			res = minDistance
		}
		prevDistance = minDistance
	}
	return res
}

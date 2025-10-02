package p1371

func findTheLongestSubstring(s string) int {
	l := len(s)
	bitSums := 0
	firsBitSums := [32]int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
	res := 0
	for i := 0; i < l; i++ {
		bit := 0
		switch s[i] {
		case 'a':
			bit = 0b00001
		case 'e':
			bit = 0b00010
		case 'i':
			bit = 0b00100
		case 'o':
			bit = 0b01000
		case 'u':
			bit = 0b10000
		}
		bitSums ^= bit
		if firsBitSums[bitSums] == -1 && bitSums != 0 {
			firsBitSums[bitSums] = i
		}
		res = max(res, i-firsBitSums[bitSums])
	}
	return res
}

func findTheLongestSubstring1(s string) int {
	l := len(s)
	bitmap := map[uint8]int{
		'a': 0b00001,
		'e': 0b00010,
		'i': 0b00100,
		'o': 0b01000,
		'u': 0b10000,
	}
	bitSums := 0
	firsBitSums := [32]int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
	res := 0
	for i := 0; i < l; i++ {
		bitSums ^= bitmap[s[i]]
		if firsBitSums[bitSums] == -1 && bitSums != 0 {
			firsBitSums[bitSums] = i
		}
		res = max(res, i-firsBitSums[bitSums])
	}
	return res
}

func findTheLongestSubstring0(s string) int {
	l := len(s)
	bitmap := map[uint8]int{
		'a': 0b00001,
		'e': 0b00010,
		'i': 0b00100,
		'o': 0b01000,
		'u': 0b10000,
	}
	elements := make([]int, l+1)
	elements[0] = 0
	bitSums := 0
	firsBitSums := make(map[int]int, l+1)
	firsBitSums[bitSums] = -1
	res := 0
	cnt, maxCnt := 0, 0
	i := 0
	for ; i < l; i++ {
		bitSums ^= bitmap[s[i]]
		if v, ok := firsBitSums[bitSums]; !ok {
			firsBitSums[bitSums] = i
		} else {
			if res < i-v {
				res = i - v
			}
		}
		elements[i+1] = bitSums
	}

	//for i := 0; i < l+1; i++ {
	//	fmt.Printf("%d: %b %d\n", i, elements[i].prefBitSum, elements[i].consonantSum)
	//}

	return max(maxCnt, cnt, res)
}

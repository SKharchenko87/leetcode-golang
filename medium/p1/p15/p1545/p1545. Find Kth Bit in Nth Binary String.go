package p1545

import "fmt"

func findKthBit(n int, k int) byte {
	positionInSection := k & -k
	isInInvertedPart := ((k / positionInSection) >> 1 & 1) == 1
	originalBitIsOne := (k & 1) == 0
	if isInInvertedPart {
		if originalBitIsOne {
			return '0'
		} else {
			return '1'
		}
	} else {
		if originalBitIsOne {
			return '1'
		} else {
			return '0'
		}
	}
}

func findKthBit1(n int, k int) byte {
	k--
	if k%8 != 3 && k%8 != 7 {
		str := "011_001_"
		return str[k%8]
	}
	prev := []byte{'0'}
	for i := 1; i < n; i++ {
		tmp := make([]byte, 1<<(i+1)-1)
		j := 0
		for _, b := range prev {
			tmp[j] = b
			j++
		}
		tmp[j] = '1'
		j++
		for _, b := range reverse(invert(prev)) {
			tmp[j] = b
			j++
		}
		//fmt.Println(string(prev))
		prev = tmp
	}
	return prev[k]
}

func findKthBit0(n int, k int) byte {
	s1 := []byte{'0'}
	prev := s1
	for i := 1; i < n; i++ {
		tmp := make([]byte, 1<<(i+1)-1)
		j := 0
		for _, b := range prev {
			tmp[j] = b
			j++
		}
		tmp[j] = '1'
		j++
		for _, b := range reverse(invert(prev)) {
			tmp[j] = b
			j++
		}
		//fmt.Println(string(prev))
		prev = tmp
	}
	fmt.Println(string(prev))
	return prev[k-1]
}

func reverse(data []byte) []byte {
	res := make([]byte, len(data))
	for i := 0; i < len(data); i++ {
		res[i] = data[len(data)-i-1]
	}
	return res
}

func invert(data []byte) []byte {
	res := make([]byte, len(data))
	for i := 0; i < len(data); i++ {
		if data[i] == '0' {
			res[i] = '1'
		} else {
			res[i] = '0'
		}
	}
	return res
}

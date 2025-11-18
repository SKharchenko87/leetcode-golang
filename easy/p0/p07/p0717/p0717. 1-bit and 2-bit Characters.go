package p0717

/*

 0|0 true
 1|0 false


00|0 true
01|0 false
10|0 true
11|0 true

000|0 true
001|0 false
010|0 true
011|0 true
100|0 true
101|0 false
110|0 true
111|0 false

0000|0 true
0001|0 false
0010|0 true
0011|0 true
0100|0 true
0101|0 false
0110|0 true
0111|0 false
1000|0 true
1001|0 false
1010|0 true
1011|0 true
1100|0 true
1101|0 false
1110|0 true
1111|0 true

1. ....1|0
2.
	11111|0 -> 111__|0 -> 1____|0 => false
	01111|0 -> 011__|0 -> 0____|0 => true
	10111|0 -> 101__|0 -> false
	10011|0 -> 100__|0 -> true

*/

func isOneBitCharacter(bits []int) bool {
	i := len(bits) - 2
	for ; i > 0 && bits[i] == 1; i -= 2 {
		if bits[i-1] == 0 {
			return false
		}
	}
	if i == 0 && bits[0] == 1 {
		return false
	}
	return true
}

func isOneBitCharacter2(bits []int) bool {
	i := len(bits) - 2
	for ; i > 0; i -= 2 {
		if bits[i] == 1 {
			if bits[i-1] == 0 {
				return false
			}
			continue
		}
		return true
	}
	if i == 0 && bits[0] == 1 {
		return false
	}
	return true
}

func isOneBitCharacter1(bits []int) bool {
	l := len(bits)
	i := 0
	for ; i < l-1; i++ {
		i += bits[i]
	}
	return i != l
}

func isOneBitCharacter0(bits []int) bool {
	l := len(bits)
	i := 0
	for ; i < l-1; i++ {
		if bits[i] == 1 {
			i++
		}
	}
	return i != l
}

package p2131

func wordsToNum(str string) uint16 {
	high := uint16(str[0] - 'a')
	low := uint16(str[1] - 'a')
	return high*26 + low
}

func reversNum(num uint16) uint16 {
	high := num / 26
	low := num % 26
	return high + low*26
}

func longestPalindrome(words []string) (res int) {
	freq := make([]int, 26*26)
	for _, word := range words {
		num := wordsToNum(word)
		rNum := reversNum(num)
		if freq[rNum] > 0 {
			freq[rNum]--
			res += 4
		} else {
			freq[num]++
		}
	}
	var i uint16
	for ; i < 26; i++ {
		if freq[i*27] > 0 {
			res += 2
			return
		}
	}
	return
}

func longestPalindrome1(words []string) int {
	freq := make([]int, 26*26)
	for _, word := range words {
		num := wordsToNum(word)
		freq[num]++
	}
	flg := 0
	res := 0
	var i uint16
	for ; i < 26*26; i++ {
		if freq[i] > 0 {
			j := reversNum(i)
			if i == j {
				if freq[i]%2 == 1 {
					res += freq[i] - 1
					if flg == 0 {
						flg = 1
					}
				} else {
					res += freq[i]
				}
			} else {
				res += min(freq[i], freq[j])
			}
		}
	}
	return (res + flg) * 2
}

func longestPalindrome0(words []string) int {
	l := len(words)
	m := make(map[[2]byte]int, l)
	for i := 0; i < l; i++ {
		b := [2]byte{words[i][0], words[i][1]}
		m[b]++
	}
	flg := 0
	res := 0
	for word, cnt := range m {
		if word[0] == word[1] {
			res += cnt / 2 * 2
			if flg == 0 && cnt%2 == 1 {
				flg = 1
			}
		} else {
			x := m[[2]byte{word[1], word[0]}]
			res += min(cnt, x)
		}
	}
	return (res + flg) * 2
}

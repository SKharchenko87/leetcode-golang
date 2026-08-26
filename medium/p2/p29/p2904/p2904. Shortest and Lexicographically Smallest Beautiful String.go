package p2904

func shortestBeautifulSubstring(s string, k int) string {
	var i, j int
	counter := 0
	l, r := -1, -1
	for ; i < len(s); i++ {
		if s[i] == '1' {
			counter++
		}

		for counter > k || (counter == k && s[j] == '0') {
			if s[j] == '1' {
				counter--
			}
			j++
		}

		if counter == k {
			curLen := i + 1 - j
			bestLen := r - l
			if l == -1 || bestLen > curLen {
				l, r = j, i+1
			} else if bestLen == curLen {
				if s[l:r] > s[j:i+1] {
					l, r = j, i+1
				}
			}
		}
	}
	if l == -1 {
		return ""
	}
	return s[l:r]
}

func shortestBeautifulSubstring0(s string, k int) string {
	var i, j int
	counter := 0
	res := s + "0"
	for ; i < len(s); i++ {
		if s[i] == '1' {
			counter++
		}
		if counter >= k {
			for ; counter > k; j++ {
				if s[j] == '1' {
					counter--
				}
			}
			for ; j < i+1; j++ {
				if len(res) > i+1-j {
					res = s[j : i+1]
				} else if len(res) == i+1-j {
					res = min(res, s[j:i+1])
				}
				if s[j] == '1' {
					break
				}
			}
		}
	}
	if res == s+"0" {
		return ""
	}
	return res
}

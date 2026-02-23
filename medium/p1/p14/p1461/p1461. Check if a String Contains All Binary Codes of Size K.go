package p1461

func hasAllCodes(s string, k int) bool {
	total := 1 << k
	if len(s) < total {
		return false
	}

	found := make([]bool, total)
	count := 0

	curMask := total - 1
	current := 0

	for i := 0; i < len(s); i++ {
		current = ((current << 1) & curMask) | int(s[i]-'0')
		if i >= k-1 {
			if !found[current] {
				found[current] = true
				count++
				if count == total {
					return true
				}
			}
		}
	}
	return false
}

var mask = 0b1111_1111_1111_1111_1111

func hasAllCodes0(s string, k int) bool {
	l := len(s)
	curMask := mask >> (20 - k)
	want := 1 << k
	m := make(map[int]struct{}, want)
	i := 0
	cur := 0
	for ; i < min(k, l); i++ {
		cur <<= 1
		cur |= int(s[i] - '0')
	}
	m[cur] = struct{}{}
	for ; i < l; i++ {
		cur <<= 1
		cur |= int(s[i] - '0')
		cur &= curMask
		m[cur] = struct{}{}
		if len(m) == want {
			return true
		}
	}
	return false
}

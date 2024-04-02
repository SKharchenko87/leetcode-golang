package p0205

func isIsomorphic(s string, t string) bool {
	ls, lt := len(s), len(t)
	if ls != lt {
		return false
	}
	ms, mt := map[byte]byte{}, map[byte]byte{}
	for i := 0; i < ls; i++ {
		is, it := s[i]-'a', t[i]-'a'
		if x, ok := mt[it]; ok {
			if x != is {
				return false
			}
		} else {
			mt[it] = is
		}
		if x, ok := ms[is]; ok {
			if x != it {
				return false
			}
		} else {
			ms[is] = it
		}
	}
	return true
}

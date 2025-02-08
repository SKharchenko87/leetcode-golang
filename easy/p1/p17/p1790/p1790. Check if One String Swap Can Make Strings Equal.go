package p1790

func areAlmostEqual(s1 string, s2 string) bool {
	counts := 0
	s1diff := [2]byte{}
	s2diff := [2]byte{}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			counts++
			if counts > 2 {
				return false
			}
			s1diff[counts-1] = s1[i]
			s2diff[counts-1] = s2[i]
		}
	}
	return s1diff[0] == s2diff[1] && s1diff[1] == s2diff[0]
}

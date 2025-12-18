package p2147

const MOD = 1e9 + 7

func numberOfWays(corridor string) int {
	l := len(corridor)
	sLastIndex := -1
	s := 0
	res := 1
	i := 0
	for ; i < l && s < 2; i++ {
		if corridor[i] == 'S' {
			sLastIndex = i
			s++
		}
	}
	for ; i < l; i++ {
		if corridor[i] == 'S' {
			if s%2 == 0 {
				res = (res * (i - sLastIndex)) % MOD
			}
			sLastIndex = i
			s++
		}
	}
	if s < 2 || s%2 == 1 || l < 2 {
		return 0
	}
	return res
}

func numberOfWays2(corridor string) int {
	l := len(corridor)
	sLastIndex := -1
	s := 0
	res := 1
	i := 0
	for ; i < l && s < 2; i++ {
		if corridor[i] == 'S' {
			sLastIndex = i
			s++
		}
	}
	for ; i < l; i++ {
		if corridor[i] == 'S' {
			if s%2 == 0 {
				res = (res * (i - sLastIndex)) % MOD
			}
			sLastIndex = i
			s++
		}
	}
	if s%2 == 1 || l < 2 || s < 2 {
		return 0
	}
	return res
}

func numberOfWays1(corridor string) int {
	l := len(corridor)
	s := 0
	i := 0
	for ; i < l && s < 2; i++ {
		if corridor[i] == 'S' {
			s++
		}
	}
	if i == l {
		if s == 2 {
			return 1
		}
		return 0
	}
	p := 0
	res := 1
	s = 0
	for ; i < l; i++ {
		if corridor[i] == 'S' {
			if s == 0 {
				res = (res * (p + 1)) % MOD
			}
			s++
			if s == 2 {
				s = 0
				p = 0
			}
		} else if s == 0 {
			p++
		}
	}
	if s == 1 {
		return 0
	}
	return res
}
